package numeral

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Numeric     int
		Roman       string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to I", 2, "II"},
		{"3 gets converted to I", 3, "III"},
	}

	for _, tt := range cases {
		t.Run(tt.Description, func(t *testing.T) {
			got := ConvertToRoman(tt.Numeric)
			if got != tt.Roman {
				t.Errorf("got %q want %q", got, tt.Roman)
			}
		})
	}
}
