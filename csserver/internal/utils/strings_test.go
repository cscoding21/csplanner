package utils

import "testing"

func TestGetTruncatedText(t *testing.T) {
	testCases := []struct {
		have   string
		want   string
		length int
	}{
		{have: "abc", want: "abc", length: 5},
		{have: "abcdef", want: "abcde...", length: 5},
		{have: "The quick brown fox jumped over the lazy dog", want: "The quick brown...", length: 16},
	}

	for _, tc := range testCases {
		result := GetTruncatedText(tc.have, tc.length)

		if result != tc.want {
			t.Errorf("Error in GetTruncatedText:  \nhave: %s\nwant: %s\ngot: %s", tc.have, tc.want, result)
		}
	}
}
