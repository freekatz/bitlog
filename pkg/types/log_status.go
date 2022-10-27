package types

import "github.com/1uvu/bitlog/pkg/common"

type StatusLog struct {
	// status detail
	StatusRaw RawLog

	// resolver
	ID                common.ID
	PrevStatusLog     *StatusLog
	NextStatusLog     *StatusLog
	RelevantEventLogs []*EventLog
	RelevantChangeLog *ChangeLog
}

func (statusLog *StatusLog) String() string {
	return ""
}
