package wordcount

import (
	"bytes"
	"strconv"
	"testing"
)

type expectedData struct {
	word  string
	count uint16
}

var text string = "hoge fuga fuga"
var expecteddata = []expectedData{
	{"hoge", 1},
	{"fuga", 2},
}

func TestAddWords(t *testing.T) {
	words := make(Words, 0, 0)
	words.AddWords(text)

	for i := range expecteddata {
		expected := &expecteddata[i]
		if words[i].Name != expected.word || words[i].Count != expected.count {
			t.Errorf("Expected %v: %d, but %v: %d", expected.word, expected.count, words[i].Name, words[i].Count)
		}
	}
}

func TestString(t *testing.T) {
	buf := &bytes.Buffer{}
	words := make(Words, 0, 0)
	words.AddWords(text)
	words.String(buf)
	outputString := buf.String()
	var expected string
	for i := range expecteddata {
		e := &expecteddata[i]
		expected += e.word + " " + strconv.FormatUint(uint64(e.count), 10) + "\n"
	}

	if outputString != expected {
		t.Errorf("Expected:\n%v\nbut:\n%v", expected, outputString)
	}
}

func TestWordCount(t *testing.T) {
	wordcount := wordCount(text)

	for i := range expecteddata {
		expected := &expecteddata[i]
		if wordcount[expected.word] != expected.count {
			t.Log("test case: ", expected.word)
			t.Errorf("Expected %d, but %d", expected.count, wordcount[expected.word])
		}
	}
}
