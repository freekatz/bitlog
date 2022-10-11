package rpcclient

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

type RPCPool struct {
	sync.RWMutex

	option             *RPCOption
	rookie, live, idle map[int]*RPCConn
	connIdx            int // a incremental id
	connNum, poolSize  int // conn num and pool size, conn num may > pool size
}

func NewRPCPool(option *RPCOption) (*RPCPool, error) {
	rookie := make(map[int]*RPCConn)
	for i := 0; i < option.PoolSize; i++ {
		c, err := NewRPCConn(i, option)
		if err != nil {
			return nil, err
		}
		rookie[i] = c
	}

	rpcPool := &RPCPool{
		option:   option,
		rookie:   rookie,
		live:     make(map[int]*RPCConn),
		idle:     make(map[int]*RPCConn),
		connIdx:  option.PoolSize,
		connNum:  option.PoolSize,
		poolSize: option.PoolSize,
	}
	return rpcPool, nil
}

func (p *RPCPool) String() string {
	return fmt.Sprintf("{rookie:%d, live:%d, idle:%d, connIdx:%d, connNum:%d, poolSize:%d}", len(p.rookie), len(p.live), len(p.idle), p.connIdx, p.connNum, p.poolSize)
}

func (p *RPCPool) selectConn() (*RPCConn, error) {
	p.RLock()
	defer p.RUnlock()

	m, n := len(p.idle), len(p.rookie)
	if m == 0 && n == 0 {
		// no idle and rookie conn
		c, err := p.createConn()
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	var key = 0
	randKey := func(pool map[int]*RPCConn) int {
		var rk int
		r := rand.Intn(len(pool))
		i := 0
		for k := range pool {
			i++
			rk = k
			if i > r {
				break
			}
		}
		return rk
	}

	// if m > 0, there has probability of 80% to select conn from idle
	randp := rand.Intn(100)
	var (
		conn *RPCConn
		ok   bool
	)
	if m > 0 && randp < 80 {
		key = randKey(p.idle)
		conn, ok = p.idle[key]
	} else {
		key = randKey(p.rookie)
		conn, ok = p.rookie[key]
	}

	if !ok {
		log.Println(key)
		panic("conn is nil")
	}
	return conn, nil
}

func (p *RPCPool) createConn() (*RPCConn, error) {
	conn, err := NewRPCConn(p.connIdx+1, p.option)
	if err != nil {
		return nil, err
	}
	p.connIdx++
	p.connNum++
	p.rookie[conn.ID] = conn
	return conn, nil
}

func (p *RPCPool) releaseConn() {
	m, n := len(p.idle), len(p.rookie)
	needRelease := m != 0 || n != 0                     // has idle or rookie conn
	needRelease = needRelease && p.connNum > p.poolSize // curr conn num > pool size
	needRelease = needRelease && p.poolSize/m < 5       // idle conn > pool size/5
	if !needRelease {
		return
	}
	releaseCount := m / 2 // c from idle
	for i := 0; i < releaseCount; i++ {
		// 1 select
		conn, err := p.selectConn()
		if err != nil {
			// dont need to handle error, just return
			return
		}
		// 3 release from map
		delete(p.idle, conn.ID)
		p.connNum--
		// 4 close conn
		conn.Shutdown()
	}
}

func (p *RPCPool) switchConnStatus(conn *RPCConn, newStatus RPCConnStatus) {
	// 1 update types
	oldStatus := conn.SwitchStatus(newStatus)
	// 2 update old set
	switch oldStatus {
	case Rookie:
		delete(p.rookie, conn.ID)
	case Live:
		delete(p.live, conn.ID)
	case Idle:
		delete(p.idle, conn.ID)
	}
	// 3 update new set
	switch newStatus {
	case Rookie:
		p.rookie[conn.ID] = conn
	case Live:
		p.live[conn.ID] = conn
	case Idle:
		p.idle[conn.ID] = conn
	}
}
