package types

import "github.com/1uvu/bitlog/pkg/common"

type ChangeLog struct {
	// change detail
	Change RawLog

	// resolver
	ID                 common.ID
	PrevChangeLog      *ChangeLog
	NextChangeLog      *ChangeLog
	RelevantStatusLogs []*StatusLog
}

func (changeLog *ChangeLog) String() string {
	return ""
}
