package easy

import "strings"

func LengthOfLastWord(s string) int {
	words := strings.Fields(s)

	l := len(words[len(words)-1])
	return l
}
