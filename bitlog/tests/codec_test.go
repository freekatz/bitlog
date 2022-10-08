package tests

import (
	"github.com/1uvu/bitlog/pkg/common"
	"testing"
)

func TestCodecBase64(t *testing.T) {
	codec := common.CodecBase64{}
	s := codec.Encode([]byte("test codec"))
	d, err := codec.Decode(s)
	t.Log(s, string(d), err)
}
