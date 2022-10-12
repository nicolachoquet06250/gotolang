package consts

import (
	"gotolang/types"
	"gotolang/utils"
)

func Interpret(arr *[][]string, line, min, max int) (*types.InterpretedConst[string], int) {
	s := (*arr)[line][min:max]

	var name string

	if s[1] != "" {
		name = s[1]
	}

	if s[3] == `"` && s[4] != "" {
		var v = s[4]
		return &(types.InterpretedConst[string]{
			Name:  name,
			Type:  types.ICT_STRING,
			Value: v,
		}), max
	} else if s[3] != `"` {
		var v = s[3]
		return &(types.InterpretedConst[string]{
			Name:  name,
			Type:  types.ICT_INT,
			Value: v,
		}), max
	}

	return nil, 0
}

func Create(t *types.InterpretedConst[string]) (lastInstruction *types.Instruction[string]) {
	_type := types.CastToInstructionValueType(types.ICT_STRING)

	if utils.MatchRegex(`(?m)[0-9]+`, t.Value) {
		_type = types.CastToInstructionValueType(types.ICT_INT)
	}

	lastInstruction = types.NewConst(
		t.Name,
		types.AFFECTATION,
		t.Value,
		_type,
	)

	return
}
