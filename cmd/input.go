package cmd

import (
	"io/ioutil"
	"os"
	"runtime"
)

func getInput(args string) (string, error) {
	input, err := getArgs(args)
	if err != nil {
		return "", err
	}
	if len(input) == 0 {
		input, err = getPipe()
		if err != nil {
			return "", err
		}
	}
	return input, nil
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func getPipe() (string, error) {
	stdin := os.Stdin
	stat, err := stdin.Stat()
	if runtime.GOOS == "windows" {
		if err == nil {
			return getStdin(stdin)
		}
	} else {
		if err != nil {
			return "", err
		}
		mode := stat.Mode()
		if (mode & os.ModeNamedPipe != 0) || mode.IsRegular() {
			return getStdin(stdin)
		}
	}
	return "" ,nil
}

func getStdin(stdin *os.File) (string, error) {
	bytes, err := ioutil.ReadAll(stdin)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func getArgs(args string) (string, error) {
	if Exists(args) {
		content, err := ioutil.ReadFile(args)
		if err != nil {
			return "", err
		}
		return string(content), nil
	}
	return args, nil
}
