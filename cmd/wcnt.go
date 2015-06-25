package cmd

import (
	"log"
	"sort"
	"strings"

	"github.com/codegangsta/cli"
)

func Wcnt(c *cli.Context) {
	words := make(Words, 0, 0)
	ign := c.String("ignorechar")
	del := c.String("delimiter")
	text, err := getInput(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}

	text = ignoreChar(text, ign)
	text = addDelimiter(text, del)
	if c.Bool("ignorecase") {
		text = strings.ToLower(text)
	}
	words.AddWords(text)

	if c.Bool("reverse") {
		sort.Sort(words)
	} else {
		sort.Sort(sort.Reverse(words))
	}
	words.String()
}
