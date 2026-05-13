package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	pattern := `^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`
		re := regexp.MustCompile(pattern)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	pattern := "<[~*=-]*>"
        re := regexp.MustCompile(pattern)
        return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	occurences := 0
	for _, v := range lines {
		pattern := `.*".*password.*".*`
		re := regexp.MustCompile(pattern)
		if re.MatchString(strings.ToLower(v)) {
			occurences += 1
		}
	}
	return occurences
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line[0-9]*`
	re := regexp.MustCompile(pattern)
	fixedText := re.ReplaceAllString(text, "")
	return fixedText
}

func TagWithUserName(lines []string) []string {
	// AI
	var userRe = regexp.MustCompile(`User\s+(\S+)`)
	out := make([]string, len(lines))
	for i, line := range lines {
		m := userRe.FindStringSubmatch(line)
		if m == nil {
			out[i] = line
			continue
		}
		out[i] = "[USR] " + m[1] + " " + line
	}
	return out
}
