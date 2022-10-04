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

type Symbol string

var (
	EQUAL             Symbol = "="
	OEPN_BRACKET             = "{"
	CLOSE_BRACKET            = "}"
	OEPN_HOOK                = "["
	CLOSE_HOOK               = "]"
	OEPN_PARENTHESIS         = "("
	CLOSE_PARENTHESIS        = ")"
)

type SymbolsComparaisons map[Symbol]Action

var symbols = SymbolsComparaisons{
	EQUAL: AFFECTATION,
}

type Instruction[T comparable] struct {
	name    string
	action  Action
	value   T
	content *Instruction[T]
}

func NewConst[T comparable](name string, action Action, value T) *Instruction[T] {
	return &Instruction[T]{
		name:    name,
		action:  action,
		value:   value,
		content: nil,
	}
}

func NewFunctionCall[T comparable](name string, content *Instruction[T]) *Instruction[T] {
	return &Instruction[T]{name: name, content: content}
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

	/*var t []Instruction[interface{ int | string }]

	t = append(t, NewConst("var", AFFECTATION, 12))

	var lastConst = t[0].(Instruction[int])
	t = append(t, NewFunctionCall("print", &lastConst))

	for _, line := range splitCode {
		for _, row := range line {
			println(row)
		}
	}

	for _, e := range t {
		println(e)
	}*/
}
