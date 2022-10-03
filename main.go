package main

import (
	"os"
	"path/filepath"
	"strings"
)

type Action string

var (
	AFFECTATION Action = "="
)

type Const[T interface{ int | string }] struct {
	name   string
	action Action
	value  T
}

func main() {
	var err error

	file := os.Args[1]
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	completePathFile := filepath.FromSlash(dir + "/" + file)

	data, err := os.ReadFile(completePathFile)
	if err != nil {
		panic(err)
	}

	l := strings.Split(string(data), LinesBreak)
	var splitCode [][]string

	for _, line := range l {
		splitLine := strings.Split(line, " ")
		var t []string
		for _, row := range splitLine {
			t = append(t, row)
		}
		splitCode = append(splitCode, t)
	}

	for _, line := range splitCode {
		for _, row := range line {
			println(row)
		}
	}

	// println(string(data), LinesBreak, l[0])
}
