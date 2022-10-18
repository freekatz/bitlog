package types

import "github.com/1uvu/bitlog/pkg/common"

type EventLog struct {
	// event detail
	EventRaw RawLog

	// resolver
	ID                common.ID
	PrevEventLog      *EventLog
	NextEventLog      *EventLog
	RelevantResultLog *ResultLog
	RelevantStatusLog *StatusLog
}

func (eventLog *EventLog) String() string {
	return ""
}
