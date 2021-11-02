package util

import (
	"regexp"
	"strings"
)

func GenAliasFor(str string) string {
	lstr := strings.ToLower(str)
	re := regexp.MustCompile(`([0-9a-z]+)`)
	slice := re.FindAllString(lstr, -1)
	return strings.Join(slice, "-")
}
