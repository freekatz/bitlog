package rpcclient

type RPCClient struct {
	*RPCPool

	Option *RPCOption
}

func NewRPCClient(option *RPCOption) (*RPCClient, error) {
	rpcPool, err := NewRPCPool(option)
	if err != nil {
		return nil, err
	}
	return &RPCClient{
		Option:  option,
		RPCPool: rpcPool,
	}, nil
}
