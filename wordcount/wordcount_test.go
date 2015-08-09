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
	wc := make(WordCounters, 0, 0)
	wc.AddWords(text)

	for i := range expecteddata {
		expected := &expecteddata[i]
		if wc[i].Word != expected.word || wc[i].Count != expected.count {
			t.Errorf("Expected %v: %d, but %v: %d", expected.word, expected.count, wc[i].Word, wc[i].Count)
		}
	}
}

func TestString(t *testing.T) {
	buf := &bytes.Buffer{}
	wc := make(WordCounters, 0, 0)
	wc.AddWords(text)
	wc.String(buf)
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

func TestCountWord(t *testing.T) {
	cw := countWord(text)

	for i := range expecteddata {
		expected := &expecteddata[i]
		if cw[expected.word] != expected.count {
			t.Log("test case: ", expected.word)
			t.Errorf("Expected %d, but %d", expected.count, cw[expected.word])
		}
	}
}
