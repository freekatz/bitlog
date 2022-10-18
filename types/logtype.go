package types

type RawLogType string

func (t RawLogType) String() string {
	return string(t)
}

// ChangeType
const (
	ChangeTypeTx    = RawLogType("change_tx")
	ChangeTypeBlock = RawLogType("change_block")

	ChangeTypeChain                 = RawLogType("change_chain")
	ChangeTypeChain_Main_Create     = RawLogType("change_chain_main_create")
	ChangeTypeChain_Main_Remove     = RawLogType("change_chain_main_remove")
	ChangeTypeChain_Main_Extend     = RawLogType("change_chain_main_extend")
	ChangeTypeChain_Main_Reduce     = RawLogType("change_chain_main_reduce")
	ChangeTypeChain_Main_Fork       = RawLogType("change_chain_main_fork")
	ChangeTypeChain_Main_Reorganize = RawLogType("change_chain_main_reorganize")

	ChangeTypeChain_Side_Create     = RawLogType("change_chain_side_create")
	ChangeTypeChain_Side_Remove     = RawLogType("change_chain_side_remove")
	ChangeTypeChain_Side_Extend     = RawLogType("change_chain_side_extend")
	ChangeTypeChain_Side_Reduce     = RawLogType("change_chain_side_reduce")
	ChangeTypeChain_Side_Fork       = RawLogType("change_chain_side_fork")
	ChangeTypeChain_Side_Reorganize = RawLogType("change_chain_side_reorganize")

	ChangeTypeNetwork = RawLogType("change_network")

	ChangeTypeUnknown = RawLogType("change_unknown")
)

// StatusType
const (
	StatueTypeChain = RawLogType("status_chain")

	StatueTypeNetwork = RawLogType("status_network")

	StatusTypeUnknown = RawLogType("status_unknown")
)

// EventType for ChangeTypeBlock
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

// EventType for ChangeTypeChain
const (
	// EventTypeChain chain
	EventTypeChain = RawLogType("event_chain")
)

// EventType for ChangeTypeNetwork
const (
	EventTypeNetwork = RawLogType("event_network")
	EventTypeUnknown = RawLogType("event_unknown")
)

// ResultType for ChangeTypeBlock
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

// ResultType for ChangeTypeChain
const (
	// ResultTypeChain chain
	ResultTypeChain = RawLogType("result_chain")
)

// ResultType for ChangeTypeNetwork
const (
	ResultTypeNetwork = RawLogType("result_network")
	ResultTypeUnknown = RawLogType("result_unknown")
)
