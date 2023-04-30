package unification

import "strings"

func CompareStrings(a string, b string, flagi bool) bool {
	if flagi {
		return strings.EqualFold(a, b)
	} else {
		return a == b
	}
}
