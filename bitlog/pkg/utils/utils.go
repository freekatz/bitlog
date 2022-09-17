package utils

import (
	"strings"
)

const DefaultConfigType = "yaml"

func GetConfigType(confPath string) string {
	if strings.HasSuffix(confPath, "yaml") {
		return "yaml"
	}

	return DefaultConfigType
}
