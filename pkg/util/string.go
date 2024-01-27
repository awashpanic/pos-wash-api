package util

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func RemoveSpecialCharacters(str string) string {
	str = strings.ToLower(str)

	// replace multiple trailing hyphens with a single hyphen
	reg := regexp.MustCompile(`(\s*-|\s+-\s*|\s+-)\s*`)
	str = reg.ReplaceAllString(str, "-")

	// replace all non-alphabet characters with hyphens
	reg = regexp.MustCompile(`[^a-zA-Z]+`)
	str = reg.ReplaceAllString(str, "-")

	// remove leading and trailing hyphens
	str = strings.Trim(str, "-")

	// replace consecutive hyphens with a single hyphen
	reg = regexp.MustCompile("-+")
	str = reg.ReplaceAllString(str, "-")

	return str
}

func Slugify(str string) string {
	str = strings.ToLower(str)

	reg := regexp.MustCompile(`(\s*-|\s+-\s*|\s+-)\s*`)
	str = reg.ReplaceAllString(str, "-")

	reg = regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	str = reg.ReplaceAllString(str, "-")
	str = strings.Trim(str, "-")

	reg = regexp.MustCompile("-+")
	str = reg.ReplaceAllString(str, "-")

	num, err := GenerateRandomNumber(5)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("%s-%d", str, num)
}
