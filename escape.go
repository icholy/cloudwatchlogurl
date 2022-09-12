package cloudwatchlogurl

import (
	"fmt"
	"net/url"
	"strings"
)

var unescape = strings.NewReplacer(
	"+", "%20",
	"%21", "!",
	"%27", "'",
	"%28", "(",
	"%29", ")",
	"%2A", "*",
)

// Go's url.QueryEscape differes from javascript's encodeURIComponent.
// It does follow the spec, but it causes issues with the aws urls.
func QueryEscape(s string) string {
	s = url.QueryEscape(s)
	return unescape.Replace(s)
}

// EcmaEscape is an implementation of javascript's deprecated escape() function
func EcmaEscape(s string) string {
	var b strings.Builder
	for _, r := range s {
		if ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || ('0' <= r && r <= '9') ||
			r == '@' || r == '*' || r == '_' || r == '+' || r == '-' || r == '.' || r == '/' {
			b.WriteRune(r)
			continue
		}
		if r >= 256 {
			fmt.Fprintf(&b, "%%%04X", r)
		} else {
			fmt.Fprintf(&b, "%%%02X", r)
		}
	}
	return b.String()
}

// FragmentEscape escapes a string for use in aws url fragments
func FragmentEscape(s string) string {
	s = QueryEscape(s)
	s = EcmaEscape(s)
	return strings.ReplaceAll(s, "%", "$")
}
