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

	_, err = os.Stat(args)
	if err != nil {
		return args, nil
	}
	content, err := ioutil.ReadFile(args)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
