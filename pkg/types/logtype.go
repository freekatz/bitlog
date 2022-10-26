package types

// TODO 完善，确定当前阶段需要的 type，边改 btcd 边确定
// TODO 根据 type 确定 status 的属性，以及其他 log 的详细内容
// TODO 完善，后面每个 type 可能都对应一种结构体从 raw 解析得到

type RawLogType string

func (t RawLogType) String() string {
	return string(t)
}

// ChangeType
const (
	ChangeTypeTx    = RawLogType("change_tx")
	ChangeTypeBlock = RawLogType("change_block")

	// ChangeTypeChain 目前只关注 change chain，每一种 change 对应独立的日志文件
	ChangeTypeChain            = RawLogType("change_chain")
	ChangeTypeChain_Create     = RawLogType("change_chain_create")
	ChangeTypeChain_Remove     = RawLogType("change_chain_remove")
	ChangeTypeChain_Extend     = RawLogType("change_chain_extend")
	ChangeTypeChain_Reduce     = RawLogType("change_chain_reduce")
	ChangeTypeChain_Fork       = RawLogType("change_chain_fork")
	ChangeTypeChain_Reorganize = RawLogType("change_chain_reorganize")

	ChangeTypeNetwork = RawLogType("change_network")

	ChangeTypeUnknown = RawLogType("change_unknown")
)

// StatusType
const (
	StatueTypeTx = RawLogType("status_tx")

	StatueTypeBlock = RawLogType("status_block")

	StatueTypeChain = RawLogType("status_chain")

	StatueTypeNetwork = RawLogType("status_network")

	StatusTypeUnknown = RawLogType("status_unknown")
)

// EventType for StatusTypeTx
const (
	// EventTypeTx tx
	EventTypeTx = RawLogType("event_tx")
)

// EventType for StatusTypeBlock
const (
	// EventTypeBlock_Arrival arrival
	EventTypeBlock_Arrival = RawLogType("event_block_arrival")

	// EventTypeBlock_Verify verify
	EventTypeBlock_Verify = RawLogType("event_block_verify")

	// EventTypeBlock_Genesis genesis
	EventTypeBlock_Genesis = RawLogType("event_block_genesis")

	// EventTypeBlock_Connect connect
	EventTypeBlock_Connect = RawLogType("event_block_connect")

	// EventTypeBlock_Orphan orphan
	EventTypeBlock_Orphan = RawLogType("event_block_orphan")

	// EventTypeBlock_Stale stale
	EventTypeBlock_Stale = RawLogType("event_block_stale")

	// EventTypeBlock_Disconnect disconnect
	EventTypeBlock_Disconnect = RawLogType("event_block_disconnect")
)

// EventType for StatusTypeChain
const (
	// EventTypeChain chain
	EventTypeChain = RawLogType("event_chain")
)

// EventType for StatusTypeNetwork
const (
	EventTypeNetwork = RawLogType("event_network")
	EventTypeUnknown = RawLogType("event_unknown")
)

// ResultType for EventTypeTx
const (
	// ResultTypeTx tx
	ResultTypeTx = RawLogType("result_tx")
)

// ResultType for EventTypeBlock
const (
	// ResultTypeBlock_Verify verify

	ResultTypeBlock_VerifySuccess = RawLogType("result_block_verify_success")
	ResultTypeBlock_VerifyFailed  = RawLogType("result_block_verify_failed")

	// ResultTypeBlock_Connect connect

	ResultTypeBlock_Connect_MainChain    = RawLogType("result_block_connect_mainchain")     // extending mainchain
	ResultTypeBlock_Connect_SideChain    = RawLogType("result_block_connect_sidechain")     // extending or creating a sidechain
	ResultTypeBlock_Connect_Attach       = RawLogType("result_block_connect_attach")        // if reorganize attach into mainchain
	ResultTypeBlock_Connect_Detach       = RawLogType("result_block_connect_detach")        // if reorganize detach from mainchain
	ResultTypeBlock_Connect_BecomeOrphan = RawLogType("result_block_connect_become_orphan") // become orphan when connect parent not exist
	ResultTypeBlock_Connect_BecomeStale  = RawLogType("result_block_connect_become_stale")  // become stale when detach
	ResultTypeBlock_Connect_BecomeBest   = RawLogType("result_block_connect_become_best")   // become best when connect or attach

	// ResultTypeBlock_Orphan orphan

	ResultTypeBlock_Orphan_BecomeStale   = RawLogType("result_block_orphan_become_stale")   // become stale when connect into orphan
	ResultTypeBlock_Orphan_AddIntoPool   = RawLogType("result_block_orphan_add_into_pool")  // orphan store into orphan pool
	ResultTypeBlock_Orphan_ExpireDiscard = RawLogType("result_block_orphan_expire_discard") // orphan discard when expire in orphan pool

	// ResultTypeBlock_Stale stale

	ResultTypeBlock_Stale_AddIntoPool = RawLogType("result_block_stale_add_into_pool") // stale store into orphan pool
	ResultTypeBlock_Stale_Discard     = RawLogType("result_block_stale_add_into_pool") // stale discard when detach or expire in orphan pool

	// ResultTypeBlock_Disconnect disconnect

	ResultTypeBlock_Disconnect_MainChain   = RawLogType("result_block_disconnect_mainchain")
	ResultTypeBlock_Disconnect_SideChain   = RawLogType("result_block_disconnect_sidechain")
	ResultTypeBlock_Disconnect_BecomeStale = RawLogType("result_block_disconnect_become_stale") // become stale when disconnect from mainchain or sidechain
)

// ResultType for EventTypeChain
const (
	// ResultTypeChain chain
	ResultTypeChain = RawLogType("result_chain")
)

// ResultType for EventTypeNetwork
const (
	ResultTypeNetwork = RawLogType("result_network")
	ResultTypeUnknown = RawLogType("result_unknown")
)
