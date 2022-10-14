package utils

import (
	"fmt"
	"strings"
)

const Debug = true

type DebugActionType func()

var msgHistory = make(map[string]string)

func generateMessage(comment string) string {
	var lineStart = " ____"
	var lineEnd = " ----"
	for range strings.Split(comment, "") {
		lineStart += "_"
		lineEnd += "-"
	}
	lineStart += "____"
	lineEnd += "----"

	msg := fmt.Sprintf("%s\n| -- %s -- |\n%s", lineStart, comment, lineEnd)

	if msgHistory[comment] != "" {
		return ""
	}

	msgHistory[comment] = msg

	return msg
}

func DebugAction(action DebugActionType, comment string, activate ...bool) {
	if len(activate) == 0 {
		activate = append(activate, true)
	}

	if Debug && activate[0] {
		if comment != "" {
			if msg := generateMessage(comment); msg != "" {
				println(msg)
			}
		}

		action()
	}
}
