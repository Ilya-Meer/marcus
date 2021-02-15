package main

import (
	"fmt"
	"strings"
)

func PadString(str string, numSpaces int) (withSpaces string) {
	newStr := []string{}

	newStr = append(
		newStr,
		strings.Repeat(" ", numSpaces),
		str,
		strings.Repeat(" ", numSpaces),
	)

	return fmt.Sprintf(strings.Join(newStr, ""))
}

func AddBreaks(numBreaks int) (breaks string) {
	return strings.Repeat("\n", numBreaks)
}
