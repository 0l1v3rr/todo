package util

import "strings"

func CreateUrlByTitle(t string) string {
	// the characters to remove
	charactersToRemove := "'\"+!%/=(){}[]|~ˇ$ß/\\.:;?@&#<>"

	// result
	res := ""

	// for loop from 0 to the (length of the string)-1
	for i := 0; i < len(t); i++ {
		// if the current character is a space, replace it with a -
		if t[i] == ' ' {
			res += "-"
			continue
		}

		// if the character is in the toRemove list then do nothing
		if strings.Contains(charactersToRemove, string(t[i])) {
			continue
		}

		// otherwise, append the current character to the res
		res += string(t[i])
	}

	return strings.ToLower(res)
}
