package consts

import (
	"gotolang/types"
	"regexp"
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
		}), len(s)
	} else if s[3] != `"` {
		var v = s[3]
		return &(types.InterpretedConst[string]{
			Name:  name,
			Type:  types.ICT_INT,
			Value: v,
		}), len(s)
	}

	return nil, 0
}

func Create(t *types.InterpretedConst[string]) (lastInstruction *types.Instruction[string]) {
	var re = regexp.MustCompile(`(?m)[0-9]+`)

	_type := types.CastToInstructionValueType(types.ICT_STRING)
	for range re.FindAllString(t.Value, -1) {
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
