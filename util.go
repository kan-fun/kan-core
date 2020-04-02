package core

import "strings"

func isAllWhiteChar(s string) bool {
	r := strings.TrimSpace(s) == ""
	return r
}
