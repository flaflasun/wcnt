package cmd

import "strings"

type editor struct {
	text string
}

func newEditor(input string) *editor {
	e := new(editor)
	e.text = input
	return e
}

func (e *editor) ignoreChar(ign string) {
	if len(ign) > 0 {
		igns := make([]string, 1, len(ign))
		igns = strings.Split(ign, "")
		for _, c := range igns {
			e.text = strings.Replace(e.text, c, "", -1)
		}
	}
}

func (e *editor) addDelimiter(del string) {
	if len(del) > 0 {
		delimiters := make([]string, 1, len(del))
		delimiters = strings.Split(del, "")
		for _, c := range delimiters {
			e.text = strings.Replace(e.text, c, " ", -1)
		}
	}
}

func (e *editor) ignorecase() {
	e.text = strings.ToLower(e.text)
}
