package event

import "github.com/fsnotify/fsnotify"

type CollectorEvent struct {
	EventFS   fsnotify.Event
	EventTag  CollectorEventTag // tag the event types
	ErrorCode int32
	ErrorMsg  string
}

func FromFsEvent(e fsnotify.Event) CollectorEvent {
	return CollectorEvent{EventFS: e}
}

type CollectorEventTag int8 // refer to epoll

const (
	TagLoop CollectorEventTag = iota // event has been processed by loop listener
)
