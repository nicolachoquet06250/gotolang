package syntax_interpreters

import (
	"errors"
	"gotolang/syntax_interpreters/call_func"
	"gotolang/syntax_interpreters/comments"
	"gotolang/syntax_interpreters/consts"
	"gotolang/syntax_interpreters/create_function"
	"gotolang/types"
	"gotolang/utils"
	"strings"
)

const EndLine = ";"
const Void = ""

func createInstructionArray(program *types.Program) (result []*types.Instruction[string]) {
	var (
		nbCols int
		nbRows int
	)

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
					nbCols = 0
					break
				}

				lineStarted = true
				lastCol = col

				if nbRows > 0 {
					nbRows--
					break
				}

				if nbCols > 0 {
					nbCols--
					continue
				}

				if types.CONST.Is(col) {
					var t *types.InterpretedConst[string]
					t, nbCols = consts.Interpret(program.ParsedCode, y, x, len(row))
					lastInstruction = consts.Create(t)
				} else if types.Function.Is(col) {
					var t *types.InterpretedCreatedFunc
					t, nbCols = create_function.Interpret(program.ParsedCode, y, x, len(row))
					lastInstruction = create_function.Create(t)
				} else if strings.Contains(strings.Join(row, ""), "(") {
					var t *types.InterpretedCallFunc[string]
					t, nbCols = call_func.Interpret(program.ParsedCode, y, x, len(row), result)
					lastInstruction = call_func.Create(t)
				} else if types.INLINE_COMMENT.Is(col) {
					nbRows = comments.Interpret(program.ParsedCode, y, x, len(row))
					lastInstruction = nil
				} else {
					println(errors.New("keyword " + col + " doesn't exist !").Error())
				}

				if lastInstruction != nil {
					result = append(result, lastInstruction)
					lastInstruction = new(types.Instruction[string])
				}
			}
		}

		nbCols = 0
		nbRows = 0
	}

	return
}

func Interpret(program *types.Program) {
	var result = createInstructionArray(program)

	// show all parsed instructions for debug
	for _, instruction := range result {
		switch instruction.InstructionType {
		case types.ASSIGN_CONST:
			utils.DebugAction(func() {
				println(
					"instruction 'assignation de constante':",
					instruction.Name,
					string(instruction.Action),
					instruction.Value+": "+instruction.ValueType.String(),
				)
			}, "Analyse des instructions du programme principal : const", true)
			break
		case types.CALL_FUNC:
			utils.DebugAction(func() {
				println(
					"instruction 'appel de fonction':",
					instruction.Name,
					instruction.Content.Value+": "+instruction.Content.ValueType.String(),
				)
			}, "Analyse des instructions du programme principal: func(...args)", true)
			break
		}
	}
}
