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
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}
