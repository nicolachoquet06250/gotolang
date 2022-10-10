package call_func

import (
	"gotolang/types"
)

func Interpret(arr *[][]string, line, min, max int, result []*types.Instruction[string]) (*types.InterpretedCallFunc[string], int) {
	s := (*arr)[line][min:max]

	var name string

	if s[0] != "" {
		name = s[0]
	}

	var constName string
	if s[1] == "(" && s[2] != "" {
		constName = s[2]
	}

	_const := new(types.InterpretedConst[string])
	for _, e := range result {
		if e.InstructionType.Is(types.ASSIGN_CONST) && e.Name == constName {
			_const = &(types.InterpretedConst[string]{
				Name:  constName,
				Type:  e.ValueType.InterpretedConstType,
				Value: e.Value,
			})
			break
		}
	}

	return &(types.InterpretedCallFunc[string]{
		Name:  name,
		Value: _const,
	}), len(s)
}

func Create(t *types.InterpretedCallFunc[string]) (lastInstruction *types.Instruction[string]) {
	newConst := types.NewConst(
		t.Value.Name,
		types.AFFECTATION,
		t.Value.Value,
		types.CastToInstructionValueType(t.Value.Type),
	)

	if newConst != nil {
		lastInstruction = types.NewFunctionCall[string](
			t.Name,
			newConst,
		)
	}

	return
}
