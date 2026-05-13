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
	// helped by ChatGPT
	pattern := `User\s+(\S+)` // capture subgroup \S
	var resultLines = make([]string, len(lines))
	re := regexp.MustCompile(pattern)
	for i, line := range lines {
		// find the submatch
		m := re.FindStringSubmatch(line) // find the submatch
		if m == nil {
			resultLines[i] = line
			continue
		}
		resultLines[i] = "[USR]" + " " + m[1] + " " + line // then append
	}
	return resultLines
}
