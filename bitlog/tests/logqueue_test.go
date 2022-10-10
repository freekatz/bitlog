package tests

import (
	"github.com/1uvu/bitlog/collector/logclient"
	"testing"
)

func TestLogQueue(t *testing.T) {
	q := logclient.NewLogQueue()
	t.Log(q.Size())
	for i := 0; i < 27; i++ {
		exceeded := q.Enqueue([]byte{byte(i)})
		if exceeded {
			t.Log("exceed-", i, ": ", exceeded)
		}
	}
	data := q.Pack()
	t.Log(data)
}
