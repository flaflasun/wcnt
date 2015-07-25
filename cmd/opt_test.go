package cmd

import "testing"

var input string = "Hoge Fuga"

type testData struct {
	argu, expected string
}

func TestNewEditor(t *testing.T) {
	expected := input
	edit := newEditor(input)
	if edit.text != expected {
		t.Errorf("Expected %v, but %v", expected, edit.text)
	}
}

func TestIgnoreChar(t *testing.T) {
	var testdata = []testData{
		{"", input},
		{"g", "Hoe Fua"},
		{"Ha", "oge Fug"},
	}

	for i := range testdata {
		ignchar := &testdata[i]
		edit := newEditor(input)
		edit.ignoreChar(ignchar.argu)
		if edit.text != ignchar.expected {
			t.Errorf("Expected %v, but %v", ignchar.expected, edit.text)
		}
	}
}

func TestAddDelimiter(t *testing.T) {
	var testdata = []testData{
		{"", input},
		{"g", "Ho e Fu a"},
		{"Ha", " oge Fug "},
	}

	for i := range testdata {
		adddelimiter := &testdata[i]
		edit := newEditor(input)
		edit.addDelimiter(adddelimiter.argu)
		if edit.text != adddelimiter.expected {
			t.Errorf("Expected %v, but %v", adddelimiter.expected, edit.text)
		}
	}
}

func TestIgnoreCase(t *testing.T) {
	expected := "hoge fuga"
	edit := newEditor(input)
	edit.ignoreCase()
	if edit.text != expected {
		t.Errorf("Expected %v, but %v", expected, edit.text)
	}
}
