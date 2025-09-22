package utils

import (
	"strings"
)

// StrPad: 문자열 s를 targetLen 길이가 될 때까지 padChar로 채워줌
// direction = "start" (기본) 또는 "end"
func StrPad(s string, targetLen int, padChar string, direction ...string) string {
	if len(s) >= targetLen {
		return s
	}
	if len(padChar) == 0 {
		padChar = " "
	}

	// direction 기본값은 "start"
	dir := "start"
	if len(direction) > 0 && direction[0] == "end" {
		dir = "end"
	}

	padding := strings.Repeat(padChar, targetLen-len(s))

	if dir == "start" {
		return padding + s
	}
	return s + padding
}
