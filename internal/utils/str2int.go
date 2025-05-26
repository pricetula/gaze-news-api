package utils

import (
	"strconv"
)

// Str2Int function supposed to covert a string to a number whenever possible
func Str2Int(s string) (r int) {
	r, err := strconv.Atoi(s)

	if err != nil {
		r = 0
	}

	return
}
