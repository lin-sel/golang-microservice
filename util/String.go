package util

import (
	"strings"
)

// IsEmpty Check weather string empty or Not
func IsEmpty(str string) bool {
	if len(strings.TrimSpace(str)) == 0 {
		return true
	}
	return false
}
