package utils

import (
	"gotolang/constants"
	"strings"
)

func buildPath(path ...string) string {
	return strings.Join(path, constants.Slash)
}
