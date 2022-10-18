package types

import (
	"encoding/json"
	"github.com/1uvu/bitlog/pkg/common"
)

type Log interface {
	ID() common.ID
	Raw() RawLog
	Prev() Log
	Next() Log
}

// RawLog parsing from the logs output by btcd
type RawLog struct {
	Type      RawLogType `json:"type"`
	Timestamp Timestamp  `json:"timestamp"`
	Raw       []byte     `json:"raw"`
}

func (raw *RawLog) Marshal() ([]byte, error) {
	data, err := json.Marshal(*raw)
	return data, err
}

// TODO replace interface with object and add more UnmarshalXXX methods
// TODO new design of raw
