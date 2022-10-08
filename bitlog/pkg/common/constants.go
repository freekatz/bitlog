package common

var (
	Constants map[ConstantKey]string
)

type ConstantKey int

const (
	ROOT_DIR = iota
	CLIENT_DIR
	LOG_DIR
	CONFIG_DIR
)

func initConstants(consts map[ConstantKey]string) {
	Constants = consts
}

func GetConstants(key ConstantKey) string {
	return Constants[key]
}
