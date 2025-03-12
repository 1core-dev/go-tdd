package numeral

import "strings"

// ConvertToRoman converts an Arabic number to a Roman Numeral.
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, digit := range allRomanNumerals {
		for arabic >= digit.Value {
			result.WriteString(digit.Symbol)
			arabic -= digit.Value
		}
	}

	return result.String()
}

type romanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []romanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
