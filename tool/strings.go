package tool

import (
	"strings"
)

func StringBetweenContain(s string, first string, end string) string {
	firstIndex := strings.Index(s, first)
	endIndex := strings.LastIndex(s, end)
	if firstIndex == -1 || endIndex == -1 {
		return s
	}
	if firstIndex >= endIndex+len(end) {
		return s
	}
	return s[firstIndex : endIndex+len(end)]
}

func StringBetween(s string, first string, end string) string {
	firstIndex := strings.Index(s, first)
	endIndex := strings.LastIndex(s, end)
	if firstIndex == -1 || endIndex == -1 {
		return s
	}
	if firstIndex+len(first) >= endIndex {
		return s
	}
	return s[firstIndex+len(first) : endIndex]
}
