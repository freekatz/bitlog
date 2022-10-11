package collector

import (
	"context"

	"github.com/1uvu/bitlog/collector/logclient"
	"github.com/1uvu/bitlog/collector/rpcclient"
	"github.com/1uvu/bitlog/pkg/common"
	"github.com/1uvu/bitlog/pkg/config"
)

// todo mod, move collector schedule into cmd backend
// TODO 解决 panic 和各种临时 log 输出

type CollectorClientMgr struct {
	// ctx  context.Context // TODO

	rpc *rpcclient.RPCClient
	log *logclient.LogClient
}

func NewCollectorClientMgr(ctx context.Context, conf *config.CollectorConfig) (*CollectorClientMgr, error) {
	rc, err := rpcclient.NewRPCClient(&rpcclient.RPCOption{
		Conf:     conf.RPC,
		PoolSize: 10,
	})
	if err != nil {
		return nil, err
	}
	lc, err := logclient.NewLogClient()
	if err != nil {
		return nil, err
	}
	mgr := new(CollectorClientMgr)
	mgr.rpc = rc
	mgr.log = lc
	return mgr, nil
}

func (mgr *CollectorClientMgr) ClientRPC() *rpcclient.RPCClient {
	return mgr.rpc
}

func (mgr *CollectorClientMgr) ClientLog() *logclient.LogClient {
	return mgr.log
}

type Collector struct {
	// ctx       context.Context TODO
	clientmgr *CollectorClientMgr
}

func NewCollector(_ctx context.Context, _conf *config.CollectorConfig) (*Collector, error) {
	ctx := context.Background() // TODO
	// 1 base the conf, get the clientmgr and add observer for log client, each client has its controller
	clientmgr, err := NewCollectorClientMgr(ctx, _conf)
	if err != nil {
		return nil, err
	}
	collector := new(Collector)
	// collector.ctx = _ctx
	collector.clientmgr = clientmgr
	return collector, nil
}

func (c *Collector) ClientMgr() *CollectorClientMgr {
	return c.clientmgr
}

func (c *Collector) HandlerLog(bufferSize int32, pool *common.Pool, filterStr string) (*logclient.LogHandler, error) {
	return logclient.NewLogHandler(c.ClientMgr().ClientLog(), bufferSize, pool, filterStr)
}

func (c *Collector) DefaultHandlerLog() (*logclient.LogHandler, error) {
	return c.HandlerLog(0, nil, "")
}

func (c *Collector) HandlerRPC(option *rpcclient.RPCOption) (*rpcclient.RPCHandler, error) {
	return rpcclient.NewRPCHandler(option)
}

func (c *Collector) DefaultHandlerRPC() (*rpcclient.RPCHandler, error) {
	return c.HandlerRPC(c.ClientMgr().ClientRPC().Option)
}
