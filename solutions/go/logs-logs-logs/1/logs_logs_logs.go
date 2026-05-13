package logs

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	pattern := "❗"
	re := regexp.MustCompile(pattern)
	re.MatchString(log)

	recommendationMatched, _ := regexp.MatchString("^❗.*", log)
	searchMatched, _ := regexp.MatchString(".*🔍.*", log)
	weatherMatched, _ := regexp.MatchString(".*☀.*", log)
	switch {
	// ugh
	case recommendationMatched:
		return "recommendation"
	case searchMatched && !weatherMatched && !recommendationMatched:
		return "search"
	case weatherMatched && !recommendationMatched && !searchMatched:
		return "weather"
	default:
		return "default"
	}
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
