package cmd

import (
	"strings"
)

func ignoreChar(t, ign string) string {
	if len(ign) > 0 {
		igns := make([]string, 1, len(ign))
		igns = strings.Split(ign, "")
		for _, c := range igns {
			t = strings.Replace(t, c, "", -1)
		}
	}
	return t
}

func addDelimiter(t, del string) string {
	if len(del) > 0 {
		delimiters := make([]string, 1, len(del))
		delimiters = strings.Split(del, "")
		for _, c := range delimiters {
			t = strings.Replace(t, c, " ", -1)
		}
	}
	return t
}
