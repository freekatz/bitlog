package rpcclient

type callFunc func(c *RPCConn, arg *ConnCallArg, reply *ConnCallReply)

// TODO implement the needed
var callHandleFuncs = map[string]callFunc{
	// "addnode":                handleAddNode,
	// "createrawtransaction":   handleCreateRawTransaction,
	// "debuglevel":             handleDebugLevel,
	// "decoderawtransaction":   handleDecodeRawTransaction,
	// "decodescript":           handleDecodeScript,
	// "estimatefee":            handleEstimateFee,
	// "generate":               handleGenerate,
	// "getaddednodeinfo":       handleGetAddedNodeInfo,
	"getbestblock": handleGetBestBlock,
	// "getbestblockhash":       handleGetBestBlockHash,
	// "getblock":               handleGetBlock,
	// "getblockchaininfo":      handleGetBlockChainInfo,
	// "getblockcount":          handleGetBlockCount,
	// "getblockhash":           handleGetBlockHash,
	// "getblockheader":         handleGetBlockHeader,
	// "getblocktemplate":       handleGetBlockTemplate,
	// "getcfilter":             handleGetCFilter,
	// "getcfilterheader":       handleGetCFilterHeader,
	// "getconnectioncount":     handleGetConnectionCount,
	// "getcurrentnet":          handleGetCurrentNet,
	// "getdifficulty":          handleGetDifficulty,
	// "getgenerate":            handleGetGenerate,
	// "gethashespersec":        handleGetHashesPerSec,
	// "getheaders":             handleGetHeaders,
	// "getinfo":                handleGetInfo,
	// "getmempoolinfo":         handleGetMempoolInfo,
	// "getmininginfo":          handleGetMiningInfo,
	// "getnettotals":           handleGetNetTotals,
	// "getnetworkhashps":       handleGetNetworkHashPS,
	// "getnodeaddresses":       handleGetNodeAddresses,
	// "getpeerinfo":            handleGetPeerInfo,
	// "getrawmempool":          handleGetRawMempool,
	// "getrawtransaction":      handleGetRawTransaction,
	// "gettxout":               handleGetTxOut,
	// "help":                   handleHelp,
	// "node":                   handleNode,
	// "ping":                   handlePing,
	// "searchrawtransactions":  handleSearchRawTransactions,
	// "sendrawtransaction":     handleSendRawTransaction,
	// "setgenerate":            handleSetGenerate,
	// "signmessagewithprivkey": handleSignMessageWithPrivKey,
	// "stop":                   handleStop,
	// "submitblock":            handleSubmitBlock,
	// "uptime":                 handleUptime,
	// "validateaddress":        handleValidateAddress,
	// "verifychain":            handleVerifyChain,
	// "verifymessage":          handleVerifyMessage,
	// "version":                handleVersion,
}

type ConnCallArg []interface{}

type ConnCallReply struct {
	Reply []interface{}
	Err   error
}

func handleGetBestBlock(c *RPCConn, arg *ConnCallArg, reply *ConnCallReply) {
	h, d, err := c.GetBestBlock()
	reply.Reply = []interface{}{h, d}
	reply.Err = err
}
