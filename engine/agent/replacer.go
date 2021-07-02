package agent

import (
	"chasqi/processor/random"
	"regexp"
	"strconv"
	"strings"
)

func ReplaceRandomsInString(raw string) string {
	generator := random.Generator{}
	re := regexp.MustCompile(`(?m)random::{\d*}`)
	match := re.FindAllStringSubmatch(raw, -1)
	newString := raw
	for _, v := range match {
		digRe := regexp.MustCompile(`\d*`)
		digVal := digRe.FindAllStringSubmatch(v[0], -1)
		parsedIntValue, err := strconv.Atoi(digVal[len(digVal)-2][0])
		if err != nil {
			panic(err)
		}
		newString = strings.Replace(newString, v[0], generator.GenerateRandomString(parsedIntValue), -1)
	}
	return strings.ToLower(newString)
}
