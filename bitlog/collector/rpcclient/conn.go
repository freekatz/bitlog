package rpcclient

import (
	"net"
	"strconv"
	"sync"

	"github.com/btcsuite/btcd/rpcclient"

	"github.com/1uvu/bitlog/pkg/config"
)

type (
	RPCConn struct {
		*rpcclient.Client
		sync.RWMutex

		ID     int
		Status RPCConnStatus
		// IdleTime time.Duration // conn in idle for a IdleTime
	}
	RPCOption struct {
		Conf     *config.RPConfig
		PoolSize int
		// IdleTime time.Duration
	}
	RPCConnStatus int8
)

const (
	Rookie RPCConnStatus = iota // conn in rookie
	Live                        // conn is live
	Idle                        // conn is live but idle (not shutdown)
)

func NewRPCConn(id int, option *RPCOption) (*RPCConn, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         net.JoinHostPort(option.Conf.Address, strconv.Itoa(option.Conf.Port)),
		User:         option.Conf.Username,
		Pass:         option.Conf.Password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	c, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}
	return &RPCConn{
		Client: c,
		ID:     id,
		Status: Rookie,
		// IdleTime: option.IdleTime,
	}, nil
}

func (c *RPCConn) SwitchStatus(newStatus RPCConnStatus) RPCConnStatus {
	c.Lock()
	defer c.Unlock()
	oldStatus := c.Status
	if oldStatus != newStatus {
		c.Status = newStatus
	}
	return oldStatus
}

func (c *RPCConn) Call(fcn string, arg *ConnCallArg, reply *ConnCallReply) {
	c.RLock()
	defer c.RUnlock()
	if f, ok := callHandleFuncs[fcn]; ok {
		f(c, arg, reply)
	}
}
