package common

import (
	"strings"
	"unicode"
)

func ConversionToPrometheusFormat(original string) (string, string) {
	var words []string
	l := 0
	for s := original; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, strings.ToLower(s[:l]))
	}
	return strings.Join(words, "_"), strings.Join(words, " ")
}
