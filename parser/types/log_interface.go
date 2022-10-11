package types

import "github.com/1uvu/bitlog/pkg/common"

// Log process, generate and maintain by resolver
type Log interface {
	ID() common.ID
	RawLog() RawLog
	PrevLog() Log
	NextLog() Log
}

type RawLogType string

func (t RawLogType) String() string {
	return string(t)
}

// RawLog parsing from the logs output by btcd
type RawLog struct {
	Type      RawLogType
	Timestamp Timestamp
	raw       []byte
}

func (r RawLog) Marshal(raw []byte) {
	r.raw = raw
}

func (r RawLog) Unmarshal() interface{} {
	return nil
}

// TODO replace interface with object and add more UnmarshalXXX methods
// TODO new design of raw
