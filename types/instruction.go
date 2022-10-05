package types

import "reflect"

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
	DOUBLE_QUOTE             = "\""
	SIMPLE_QUOTE             = "'"
	SEMICOLUMN               = ";"
)

func IsSymbol(symbol string) bool {
	var s = new(Symbol)

	if symbol == string(EQUAL) ||
		symbol == OEPN_BRACKET ||
		symbol == CLOSE_BRACKET ||
		symbol == OEPN_HOOK ||
		symbol == CLOSE_HOOK ||
		symbol == OEPN_PARENTHESIS ||
		symbol == CLOSE_PARENTHESIS ||
		symbol == DOUBLE_QUOTE ||
		symbol == SIMPLE_QUOTE ||
		symbol == SEMICOLUMN {
		_s := Symbol(symbol)
		s = &_s
	}

	typeName := reflect.TypeOf(s).Elem().String()
	if *s == "" {
		typeName = "nil"
	}

	return typeName == "types.Symbol"
}

type Keyword string

var (
	CONST Keyword = "const"
)

func IsKeyword(keyword string) bool {
	var s = new(Keyword)

	if keyword == string(CONST) {
		_s := Keyword(keyword)
		s = &_s
	}

	typeName := reflect.TypeOf(s).Elem().String()
	if *s == "" {
		typeName = "nil"
	}

	return typeName == "types.Keyword"
}

type Instruction[T comparable] struct {
	Name    string
	Action  Action
	Value   T
	Content *Instruction[T]
}

func NewConst[T comparable](name string, action Action, value T) *Instruction[T] {
	return &Instruction[T]{
		Name:    name,
		Action:  action,
		Value:   value,
		Content: nil,
	}
}

func NewFunctionCall[T comparable](name string, content *Instruction[T]) *Instruction[T] {
	return &Instruction[T]{
		Name:    name,
		Content: content,
	}
}

type SymbolsComparaisons map[Symbol]Action
