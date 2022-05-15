package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
    exp := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)

    return exp.MatchString(text)
}

func SplitLogLine(text string) []string {
    exp := regexp.MustCompile(`<(\*|~|=|-)*>`)

    return exp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    var count int
    exp := regexp.MustCompile(`(?i)".*password.*"`)

    for _, line := range lines {
        if exp.MatchString(line) {
            count++
        }
    }

    return count
}

func RemoveEndOfLineText(text string) string {
    exp := regexp.MustCompile(`end-of-line\d+`)

    return exp.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
    var tagged []string
    exp := regexp.MustCompile(`User\s+(\w+)`)

    for _, line := range lines {
        if exp.MatchString(line) {
            name := exp.FindStringSubmatch(line)
            newLine := fmt.Sprintf("[USR] %s %s", name[len(name) - 1], line)
            tagged = append(tagged, newLine)
        } else {
            tagged = append(tagged, line)
        }
    }

    return tagged
}
