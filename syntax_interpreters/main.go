package syntax_interpreters

import (
	"gotolang/syntax_interpreters/call_func"
	"gotolang/syntax_interpreters/consts"
	"gotolang/types"
)

const EndLine = ";"
const Void = ""

func Interpret(program *types.Program) {
	var result []*types.Instruction[string]
	var nbContinue int
	for y, row := range *program.ParsedCode {
		var lastCol string
		var lastInstruction *types.Instruction[string]
		var isVoidLine = len(row) == 1 && row[0] == Void
		if !isVoidLine {
			var lineStarted = false
			for x, col := range row {
				if !lineStarted && col == Void {
					continue
				}
				if lastCol == EndLine {
					nbContinue = 0
					break
				}
				lineStarted = true
				lastCol = col

				if nbContinue > 0 {
					nbContinue--
					continue
				}

				if types.CONST.Is(col) {
					var t *types.InterpretedConst[string]
					t, nbContinue = consts.Interpret(program.ParsedCode, y, x, len(row)-1)
					lastInstruction = consts.Create(t)
				} else {
					var t *types.InterpretedCallFunc[string]
					t, nbContinue = call_func.Interpret(program.ParsedCode, y, x, len(row)-1, result)
					lastInstruction = call_func.Create(t)
				}

				if lastInstruction != nil {
					result = append(result, lastInstruction)
					lastInstruction = new(types.Instruction[string])
				}
			}
		}
		nbContinue = 0
	}

	// show all parsed instructions for debug
	for _, instruction := range result {
		switch instruction.InstructionType {
		case types.ASSIGN_CONST:
			println(
				"instruction 'assignation de constante':",
				instruction.Name,
				string(instruction.Action),
				instruction.Value+": "+instruction.ValueType.String(),
			)
			break
		case types.CALL_FUNC:
			println(
				"instruction 'appel de fonction':",
				instruction.Name,
				instruction.Content.Value+": "+instruction.Content.ValueType.String(),
			)
			break
		}
	}
}
