package types

import "github.com/1uvu/bitlog/pkg/common"

type (
	StatusLog struct {
		// status detail
		Status RawLog

		// resolver
		ID            common.ID
		PrevStatusLog *StatusLog // last status
		NextStatusLog *StatusLog // create when status changed

		RelevantEventLogs *EventLogLinkedList
		RelevantChangeLog *ChangeLog
	}
)

// StatusType
const (
	StatueTypeChain   = RawLogType("status_chain")
	StatueTypeNetwork = RawLogType("status_network")
	StatusTypeUnknown = RawLogType("status_unknown")
)

func (statusLog *StatusLog) String() string {
	return ""
}
