package common

import "encoding/base64"

type (
	Codec interface {
		Encode([]byte) string
		Decode(string) ([]byte, error)
	}
	CodecBase64 struct {
	}
)

func (codec *CodecBase64) Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func (codec *CodecBase64) Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
