package types

import "github.com/1uvu/bitlog/pkg/common"

type (
	ChangeLog struct {
		// change detail
		Change RawLog

		// resolver
		ID            common.ID
		PrevChangeLog *ChangeLog // RelevantStatusLogs.Tail.RelevantChangeLog
		NextChangeLog *ChangeLog // create when IsValid=true

		IsValid            bool
		RelevantStatusLogs []*StatusLog
	}
)

// ChangeType
const (
	ChangeTypeTx      = RawLogType("change_tx")
	ChangeTypeBlock   = RawLogType("change_block")
	ChangeTypeChain   = RawLogType("change_chain")
	ChangeTypeNetwork = RawLogType("change_network")
	ChangeTypeUnknown = RawLogType("change_unknown")
)

func (changeLog *ChangeLog) String() string {
	return ""
}
