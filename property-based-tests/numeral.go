package numeral

import "strings"

func ConvertToRoman(number int) string {
	var result strings.Builder

	for range number {
		result.WriteString("I")
	}
	return result.String()
}
