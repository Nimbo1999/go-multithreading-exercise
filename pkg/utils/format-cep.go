package utils

import (
	"regexp"
	"strings"
)

func FormatCep(cep string) string {
	if len(cep) < 8 {
		return ""
	}
	cep = removeAlphaNumericCharacters(cep)
	if len(cep) != 8 {
		return ""
	}
	splittedString := strings.Split(cep, "")
	return strings.Join(splittedString[:5], "") + "-" + strings.Join(splittedString[5:], "")
}

func removeAlphaNumericCharacters(value string) string {
	sampleRegexp := regexp.MustCompile(`\D`)
	result := sampleRegexp.ReplaceAllString(value, "")
	return string(result)
}
