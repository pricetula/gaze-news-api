package utils

import (
	"strings"
)

// function supposed to split string to a slice of strings
func SpltStr2Slc(s string, splt string) (r []string) {
	r = strings.Split(s, splt)

	return
}
