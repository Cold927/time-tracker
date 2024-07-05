package utils

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToFormatCase(str string) string {
	format := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	format = matchAllCap.ReplaceAllString(format, "${1}_${2}")
	return strings.ToLower(format)
}
