package parsinglogfiles

import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	if len(text) >= 5 {
		re := regexp.MustCompile(`^\[TRC]|^\[ERR]|^\[DBG]|^\[INF]|^\[WRN]|^\[FTL]`)
		return re.MatchString(text)
	}
	return false
}

func SplitLogLine(text string) []string {
	if len(text) > 0 {
		fmt.Println(text)
		re := regexp.MustCompile(`<[-=~*]*>`)
		return re.Split(text, -1)
	}
	return []string{""}
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	sum := 0
	for _, line := range lines {
		if re.MatchString(line) {
			sum++
		}
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	r := []string{}
	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		r = append(r, line)
	}
	return r
}
