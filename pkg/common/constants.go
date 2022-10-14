package common

var (
	constants map[ConstantKey]string
)

type ConstantKey int

func InitConstants(consts map[ConstantKey]string) {
	constants = consts
}

func GetConstants(key ConstantKey) string {
	return constants[key]
}
