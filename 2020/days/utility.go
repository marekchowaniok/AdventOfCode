package days

import "strings"

func sliceInput(bytes []byte) []string {
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	return split
}
