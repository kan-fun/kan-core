package core

import "strings"

// IsAllWhiteChar ...
func IsAllWhiteChar(s string) bool {
	r := strings.TrimSpace(s) == ""
	return r
}
