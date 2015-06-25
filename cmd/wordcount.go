package cmd

import (
	"fmt"
	"strings"
)

type Word struct {
	Name  string
	Count uint16
}

type Words []Word

func (w *Words) AddWords(t string) {
	wordcount := wordCount(t)
	*w = make(Words, 0, len(wordcount))

	for n := range wordcount {
		word := Word{
			Name:  n,
			Count: wordcount[n],
		}
		*w = append(*w, word)
	}
}

func (w *Words) String() {
	for _, word := range *w {
		fmt.Println(word.Name, word.Count)
	}
}

func wordCount(t string) map[string]uint16 {
	m := make(map[string]uint16)
	for _, word := range strings.Fields(t) {
		m[word]++
	}
	return m
}
