package utils

import (
	"fmt"
	"strings"
)

// WrapInBackticks take a string and enclose it in back ticks
func WrapInBackticks(input string) string {
	return fmt.Sprint("`" + input + "`")
}

// GetTruncatedText returns a truncated version of the input string if it is longer than max.
func GetTruncatedText(input string, max int) string {
	if len(input) > max {
		return strings.Trim(input[:max], " ") + "..."
	}

	return input
}

// CoalesceString return the first non-nil and non-empty string in the param array
func CoalesceString(args ...*string) string {
	for _, a := range args {
		res := *a
		if len(res) > 0 {
			return res
		}
	}

	return ""
}
