package main

import (
	"fmt"
	"strings"
)

func (f *Formatter) PadString(str string, numSpaces int) {
	newStr := []string{}

	newStr = append(
		newStr,
		strings.Repeat(" ", numSpaces),
		str,
		strings.Repeat(" ", numSpaces),
	)

	f.WriteString(fmt.Sprintf(strings.Join(newStr, "")))
}

func (f *Formatter) AddBreaks(numBreaks int) {
	f.WriteString(strings.Repeat("\n", numBreaks))
}
