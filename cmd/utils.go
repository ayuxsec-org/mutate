package main

import (
	"bufio"
	"os"
	"strings"
)

func fileToSlice(file *os.File) []string {
	var out []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, strings.TrimSpace(scanner.Text()))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return out
}

func isPipedStdin() bool {
	fi := must(os.Stdin.Stat())
	return (fi.Mode() & os.ModeCharDevice) == 0
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
