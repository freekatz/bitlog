package types

import "github.com/1uvu/bitlog/pkg/common"

type ResultLog struct {
	// result detail
	ResultRaw RawLog

	// resolver
	ID               common.ID
	RelevantEventLog *EventLog
}

func (resultLog *ResultLog) String() string {
	return ""
}
