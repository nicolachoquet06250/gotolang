package utils

import (
	"strings"
)

type separatorType interface{ string | []string }

func Split[T separatorType](s string, sep T) []string {
	switch (interface{})(sep).(type) {
	case string:
		return strings.Split(s, (interface{})(sep).(string))
	case []string:
		return strings.FieldsFunc(s, func(r rune) bool {
			for _, _s := range (interface{})(sep).([]string) {
				if string(r) == _s {
					return true
				}
			}
			return false
		})
	default:
		return []string{}
	}
}
