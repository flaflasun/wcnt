package cmd

import (
	"io/ioutil"
	"os"
)

func getInput(args string) (string, error) {
	stdin := os.Stdin
	stat, err := stdin.Stat()
	if err != nil {
		return "", err
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := ioutil.ReadAll(stdin)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}

	if Exists(args) {
		content, err := ioutil.ReadFile(args)
		if err != nil {
			return "", err
		}
		return string(content), nil
	}
	return args, nil
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
