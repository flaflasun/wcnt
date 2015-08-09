package cmd

import (
	"log"
	"os"
	"sort"

	"github.com/codegangsta/cli"
	"github.com/flaflasun/wordc/wordcount"
)

func Wordc(c *cli.Context) {
	input, err := getInput(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}

	edit := newEditor(input)
	edit.ignoreChar(c.String("ignorechar"))
	edit.addDelimiter(c.String("delimiter"))
	if c.Bool("ignorecase") {
		edit.ignoreCase()
	}

	wc := make(wordcount.WordCounters, 0, 0)
	wc.AddWords(edit.text)

	if c.Bool("reverse") {
		sort.Sort(wc)
	} else {
		sort.Sort(sort.Reverse(wc))
	}
	wc.String(os.Stdout)
}
