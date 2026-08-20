package logs

import (
	"strings"
    "unicode/utf8"

)

var (
	recommendation rune = '❗'
	search         rune = '🔍'
	weather        rune = '☀'
)

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, v := range log {
		switch {
		case v == recommendation:
			return "recommendation"
		case v == search:
			return "search"
		case v == weather:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newStr := strings.Replace(log, string(oldRune), string(newRune), 999)
	return newStr
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
		return true
	}
	return false
}
