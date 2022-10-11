package rpcclient

// Put on hold

type RPCHandler struct {
	c *RPCClient
}

func NewRPCHandler(option *RPCOption) (*RPCHandler, error) {
	rpcClient, err := NewRPCClient(option)
	if err != nil {
		return nil, err
	}
	return &RPCHandler{
		c: rpcClient,
	}, nil
}

// Call
// firstly select a connection to process the rpc call, and if no rookie conn will create a new automatically,
// will update the connection status, and after call success will backtrack the connection status
func (h *RPCHandler) Call(fcn string, arg *ConnCallArg, reply *ConnCallReply) {
	conn, err := h.c.selectConn()
	if err != nil {
		reply.Err = err
		return
	}
	// conn -> live
	h.c.switchConnStatus(conn, Live)
	// call
	conn.Call(fcn, arg, reply)
	// conn -> idle
	h.c.switchConnStatus(conn, Idle)
	// try release idle conn
	h.c.releaseConn()
}

// CallAsync Call async
func (h *RPCHandler) CallAsync(fcn string, arg *ConnCallArg, replyChan chan<- *ConnCallReply) {
	var reply = &ConnCallReply{}
	go func() {
		h.Call(fcn, arg, reply)
		replyChan <- reply
	}()
}
