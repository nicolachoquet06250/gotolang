package utils

import (
	"fmt"
	"strings"
)

const Debug = true

type DebugActionType func()

func DebugAction(action DebugActionType, comment string, activate ...bool) {
	if len(activate) == 0 {
		activate = append(activate, true)
	}
	if Debug && activate[0] {
		if comment != "" {
			var lineStart = " ____"
			var lineEnd = " ----"
			for range strings.Split(comment, "") {
				lineStart += "_"
				lineEnd += "-"
			}
			lineStart += "____"
			lineEnd += "----"

			println(lineStart)
			println(fmt.Sprintf("| -- %s -- |", comment))
			println(lineEnd)
		}
		action()
	}
}
