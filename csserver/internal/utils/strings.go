package utils

import (
	"fmt"
)

// WrapInBackticks take a string and enclose it in back ticks
func WrapInBackticks(input string) string {
	return fmt.Sprint("`" + input + "`")
}
