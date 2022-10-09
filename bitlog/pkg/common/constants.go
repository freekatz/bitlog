package common

var (
	constants map[ConstantKey]string
)

type ConstantKey int

const (
	ROOT_DIR = iota
	CLIENT_DIR
	LOG_DIR
	CONFIG_DIR
)

func initConstants(consts map[ConstantKey]string) {
	constants = consts
}

func GetConstants(key ConstantKey) string {
	return constants[key]
}
