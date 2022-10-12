package utils

import (
	"regexp"
)

func MatchRegex(regex string, str string) bool {
	var re = regexp.MustCompile(regex)

	for range re.FindAllString(str, -1) {
		return true
	}
	return false
}
