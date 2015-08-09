package wordcount

import (
	"fmt"
	"io"
	"strings"
)

type WordCounter struct {
	Word  string
	Count uint16
}

type WordCounters []WordCounter

func (wc *WordCounters) AddWords(t string) {
	cw := countWord(t)
	*wc = make(WordCounters, 0, len(cw))

	for w := range cw {
		word := WordCounter{
			Word:  w,
			Count: cw[w],
		}
		*wc = append(*wc, word)
	}
}

func (wc *WordCounters) String(wr io.Writer) {
	for _, w := range *wc {
		fmt.Fprintln(wr, w.Word, w.Count)
	}
}

func countWord(t string) map[string]uint16 {
	m := make(map[string]uint16)
	for _, word := range strings.Fields(t) {
		m[word]++
	}
	return m
}
