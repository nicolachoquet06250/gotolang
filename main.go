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

	syntax_interpreters.Interpret(
		types.NewProgram(splitCode),
	)

	/*var (
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
	}*/
}
