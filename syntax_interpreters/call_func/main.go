package call_func

import "gotolang/syntax_interpreters/consts"

type InterpretedCallFuncType string

var (
	STRING InterpretedCallFuncType = "string"
	INT    InterpretedCallFuncType = "int"
	CONST  InterpretedCallFuncType = "const"
)

type InterpretedCallFunc[T comparable] struct {
	Name  string
	Type  InterpretedCallFuncType
	Value *consts.InterpretedConst[string]
}

func Interpret(arr *[][]string, line, min, max int) *InterpretedCallFunc[string] {
	s := (*arr)[line][min:max]

	var name string

	if s[0] != "" {
		name = s[0]
	}

	var constName string
	if s[1] == "(" && s[2] != "" {
		constName = s[2]
	}

	return &(InterpretedCallFunc[string]{
		Name: name,
		Type: CONST,
		Value: &(consts.InterpretedConst[string]{
			Name:  constName,
			Type:  consts.STRING,
			Value: "test",
		}),
	})
}
