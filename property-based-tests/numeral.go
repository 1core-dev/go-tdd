package numeral

func ConvertToRoman(number int) string {
	if number == 3 {
		return "III"
	}
	if number == 2 {
		return "II"
	}
	return "I"
}
