package cloudwatchlogurl

import (
	"sort"
	"strings"
)

// QueryDetails encodes key/value pairs into the format used in aws
// log insights urls.
type QueryDetails map[string][]string

// Add a value to the specified key. The quote parameter indicates if
// the provided value should be quoted.
func (q QueryDetails) Add(key, value string, quote bool) {
	escaped := QueryEscape(value)
	escaped = strings.ReplaceAll(escaped, "%", "*")
	if quote {
		escaped = "'" + escaped
	}
	q[key] = append(q[key], escaped)
}

// Encode returns the key/values as an encoded string.
func (q QueryDetails) Encode() string {
	var b strings.Builder
	b.WriteString("~(")
	// sort the keys to get a deterministic output order
	keys := make([]string, 0, len(q))
	for key := range q {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i, key := range keys {
		if i > 0 {
			b.WriteByte('~')
		}
		b.WriteString(key)
		b.WriteByte('~')
		values := q[key]
		switch len(values) {
		case 0:
		case 1:
			b.WriteString(values[0])
		default:
			b.WriteByte('(')
			for _, v := range values {
				b.WriteByte('~')
				b.WriteString(v)
			}
			b.WriteByte(')')
		}
	}
	b.WriteByte(')')
	escaped := QueryEscape("?queryDetail=" + EcmaEscape(b.String()))
	escaped = strings.ReplaceAll(escaped, "%", "$")
	return escaped
}
