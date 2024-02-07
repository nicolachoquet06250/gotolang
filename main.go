package main

import (
	"gotolang/parser"
	"gotolang/utils"
	"os"
)

func main() {
	file := os.Args[1]
	data := utils.OpenFile(file)

	var splitCode = parser.Parse(data)

	utils.DebugAction(func() {
		for _, row := range *splitCode {
			for _, col := range row {
				println(col)
			}
		}
	}, "", false)

	/*syntax_interpreters.Interpret(
		types.NewProgram(splitCode),
	)*/

	/*c := Instruction[func(...any)]{
		Type: CreateConst,
		Const: &TConst{
			Name:  "toto",
			Value: "test",
		},
	}
	fPrint := Instruction[func(args ...any)]{
		Type: CreateFunction,
		Function: &TFunction[func(args ...any)]{
			Name: "print",
			Ref: func(args ...any) {
				for _, e := range args {
					println(e.(string))
				}
			},
		},
	}
	call := Instruction[func(args ...any)]{
		Type: CallFunction,
		Call: &TCall[func(args ...any)]{
			Func: *fPrint.Function,
			Parameters: []Instruction[func(args ...any)]{
				c,
				c,
			},
		},
	}
	call.Call.Func.Ref("start created func ref")*/
}
