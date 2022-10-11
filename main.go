package main

import (
	"gotolang/parser"
	"gotolang/syntax_interpreters"
	"gotolang/types"
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
	}, false)

	syntax_interpreters.Interpret(
		types.NewProgram(splitCode),
	)
}
