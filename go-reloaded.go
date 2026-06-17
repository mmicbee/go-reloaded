package main 

import "strings"

func Capitalize(s string) string {
	if s == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}