package iteration

import "strings"

const repeatCount = 5

func Repeat(char string) string {
	var repeated string
	for range repeatCount {
		repeated += char
	}
	return repeated
}

func RepeatV2(char string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(char)
	}
	return repeated.String()
}
