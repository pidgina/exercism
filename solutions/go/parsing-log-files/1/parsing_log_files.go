package parsinglogfiles

import "regexp"

var (
	validLineRe      = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	separatorRe      = regexp.MustCompile(`<[~*=-]*>`)
	quotedPasswordRe = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	endOfLineRe      = regexp.MustCompile(`end-of-line\d+`)
	userNameRe       = regexp.MustCompile(`User +(\S+)`)
)

func IsValidLine(text string) bool {
	return validLineRe.MatchString(text)
}

func SplitLogLine(text string) []string {
	return separatorRe.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if quotedPasswordRe.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRe.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	tagged := make([]string, 0, len(lines))

	for _, line := range lines {
		match := userNameRe.FindStringSubmatch(line)
		if match == nil {
			tagged = append(tagged, line)
			continue
		}

		tagged = append(tagged, "[USR] "+match[1]+" "+line)
	}

	return tagged
}
