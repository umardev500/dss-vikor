package utils

import (
	"regexp"
	"strings"
)

func SingleLineString(s string) string {
	s = strings.ReplaceAll(s, "\t", "")    // remove tabs
	s = strings.Trim(s, " ")               // remove leading and trailing spaces
	s = strings.ReplaceAll(s, "\n", " ")   // remove new line
	s = strings.ReplaceAll(s, "--sql", "") // remove --sql
	re := regexp.MustCompile(`\s+`)        // remove multiple spaces
	output := re.ReplaceAllString(s, " ")  // replace multiple spaces
	return output
}
