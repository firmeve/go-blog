package utils

import (
	"regexp"
	"strings"
)

func StringUcFirst(str string) string {
	first := strings.ToUpper(str[0:1])
	return strings.Join([]string{first, str[1:]}, ``)
}

func StringUcWords(words []string) string {
	return strings.ReplaceAll(strings.Title(strings.Join(words,` `)),` `,``)
}

func StringSnakeCase(str string) string {
	return strings.ToLower(regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(str, "${1}_${2}"))
}
