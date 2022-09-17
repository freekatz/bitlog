package types

import "github.com/1uvu/bitlog/pkg/common"

type (
	EventLogLinkedList struct {
		Timeline   Timeline // TODO timestamp + event = timeline
		Head, Tail *EventLog
	}
	EventLog struct {
		// event detail
		EventRaw RawLog

		// resolver
		ID           common.ID
		PrevEventLog *EventLog // RelevantStatusLog.RelevantEventLogs.Tail
		NextEventLog *EventLog // create when happen a new event

		RelevantStatusLog *StatusLog
	}
)

// EventType for ChangeTypeBlock
const (
	// EventTypeBlock_Arrival arrival
	EventTypeBlock_Arrival = RawLogType("event_block_arrival")

	// EventTypeBlock_Verify verify
	EventTypeBlock_Verify        = RawLogType("event_block_verify")
	EventTypeBlock_VerifySuccess = RawLogType("event_block_verify_success")
	EventTypeBlock_VerifyFailed  = RawLogType("event_block_verify_failed")

	// EventTypeBlock_Genesis genesis
	EventTypeBlock_Genesis = RawLogType("event_block_genesis")

	// EventTypeBlock_Connect connect
	EventTypeBlock_Connect              = RawLogType("event_block_connect")
	EventTypeBlock_Connect_MainChain    = RawLogType("event_block_connect_mainchain")     // extending mainchain
	EventTypeBlock_Connect_SideChain    = RawLogType("event_block_connect_sidechain")     // extending or creating a sidechain
	EventTypeBlock_Connect_Attach       = RawLogType("event_block_connect_attach")        // if reorganize attach into mainchain
	EventTypeBlock_Connect_Detach       = RawLogType("event_block_connect_detach")        // if reorganize detach from mainchain
	EventTypeBlock_Connect_BecomeOrphan = RawLogType("event_block_connect_become_orphan") // become orphan when connect parent not exist
	EventTypeBlock_Connect_BecomeStale  = RawLogType("event_block_connect_become_stale")  // become stale when detach
	EventTypeBlock_Connect_BecomeBest   = RawLogType("event_block_connect_become_best")   // become best when connect or attach

	// EventTypeBlock_Orphan orphan
	EventTypeBlock_Orphan               = RawLogType("event_block_orphan")
	EventTypeBlock_Orphan_BecomeStale   = RawLogType("event_block_orphan_become_stale")   // become stale when connect into orphan
	EventTypeBlock_Orphan_AddIntoPool   = RawLogType("event_block_orphan_add_into_pool")  // orphan store into orphan pool
	EventTypeBlock_Orphan_ExpireDiscard = RawLogType("event_block_orphan_expire_discard") // orphan discard when expire in orphan pool

	// EventTypeBlock_Stale stale
	EventTypeBlock_Stale             = RawLogType("event_block_stale")
	EventTypeBlock_Stale_AddIntoPool = RawLogType("event_block_stale_add_into_pool") // stale store into orphan pool
	EventTypeBlock_Stale_Discard     = RawLogType("event_block_stale_add_into_pool") // stale discard when detach or expire in orphan pool

	// EventTypeBlock_Disconnect disconnect
	EventTypeBlock_Disconnect             = RawLogType("event_block_disconnect")
	EventTypeBlock_Disconnect_MainChain   = RawLogType("event_block_disconnect_mainchain")
	EventTypeBlock_Disconnect_SideChain   = RawLogType("event_block_disconnect_sidechain")
	EventTypeBlock_Disconnect_BecomeStale = RawLogType("event_block_disconnect_become_stale") // become stale when disconnect from mainchain or sidechain
)

// EventType for ChangeTypeChain
const (
	// EventTypeChain_MainChain mainchain
	EventTypeChain_MainChain            = RawLogType("event_chain_mainchain")
	EventTypeChain_MainChain_Create     = RawLogType("event_chain_mainchain_create")     // genesis
	EventTypeChain_MainChain_Extend     = RawLogType("event_chain_mainchain_extend")     // connect
	EventTypeChain_MainChain_Reduce     = RawLogType("event_chain_mainchain_reduce")     // disconnect
	EventTypeChain_MainChain_Reorganize = RawLogType("event_chain_mainchain_reorganize") // attach and detach

	// EventTypeChain_SideChain sidechain
	EventTypeChain_SideChain            = RawLogType("event_chain_sidechain")
	EventTypeChain_SideChain_Create     = RawLogType("event_chain_sidechain_create")     // connect
	EventTypeChain_SideChain_Extend     = RawLogType("event_chain_sidechain_extend")     // connect
	EventTypeChain_SideChain_Reduce     = RawLogType("event_chain_sidechain_reduce")     // disconnect
	EventTypeChain_SideChain_Reorganize = RawLogType("event_chain_sidechain_reorganize") // attach and detach
)

// EventType for ChangeTypeNetwork
const (
	EventTypeNetwork = RawLogType("event_network")
	EventTypeUnknown = RawLogType("event_unknown")
)

func (eventLog *EventLog) String() string {
	return ""
}
