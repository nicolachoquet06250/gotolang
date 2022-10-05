package main

import (
	"fmt"
	"gotolang/parser"
	"gotolang/syntax_interpreters/call_func"
	"gotolang/syntax_interpreters/consts"
	"gotolang/types"
	"os"
)

var symbols = types.SymbolsComparaisons{
	types.EQUAL: types.AFFECTATION,
}

func main() {
	file := os.Args[1]
	data := openFile(file)

	var splitCode = parser.Parse(data)

	/*for _, row := range *splitCode {
		for _, col := range row {
			println(fmt.Sprintf(`"%s"`, col))
		}
		println("-------------------------------")
	}*/

	var (
		minIndex = -1
		maxIndex = -1
	)
	for i, row := range *splitCode {
		var lastInstructionKeyword types.Keyword
		for j, col := range row {
			if types.IsKeyword(col) && col == string(types.CONST) {
				minIndex = j
				lastInstructionKeyword = types.CONST
			} else if types.IsSymbol(col) && col == types.SEMICOLUMN {
				maxIndex = j
				if lastInstructionKeyword == types.CONST {
					interpreted := consts.Interpret(splitCode, i, minIndex, maxIndex)

					var value string
					if interpreted.Type == consts.STRING {
						value += `"`
					}
					value += interpreted.Value
					if interpreted.Type == consts.STRING {
						value += `"`
					}

					println(fmt.Sprintf(`%s=%s`, interpreted.Name, value))
				} else {
					interpreted := call_func.Interpret(splitCode, i, minIndex, maxIndex)

					var value string
					if interpreted.Value.Type == consts.STRING {
						value += `"`
					}
					value += interpreted.Value.Value
					if interpreted.Value.Type == consts.STRING {
						value += `"`
					}

					println(fmt.Sprintf(`%s(%s)`, interpreted.Name, value))
				}
				minIndex = 0
			}
		}
	}
}
