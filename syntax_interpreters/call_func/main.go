package call_func

import (
	"fmt"
	"gotolang/types"
	"gotolang/utils"
)

func handleError(err error) {
	println(err.Error())
}

func getConstName[F types.Func](result []*types.Instruction[F], constName string) *types.InterpretedConst {
	c := new(types.InterpretedConst)

	for _, e := range result {
		if e.Type.Is(types.CreateConst) && e.Const.Name == constName {
			c = utils.New[types.InterpretedConst](utils.PropertiesAny{
				{
					Key:   "Name",
					Value: constName,
				},
				{
					Key:   "Type",
					Value: fmt.Sprintf("%T", e.Const.Value),
				},
				{
					Key:   "Value",
					Value: e.Const.Value,
				},
			}, handleError)
		}
	}
	return c
}

func Interpret[F types.Func](arr *[][]string, line, min, max int, result []*types.Instruction[F]) (*types.InterpretedCallFunc, int) {
	s := (*arr)[line][min:max]

	var name string

	if s[0] != "" {
		name = s[0]
	}

	var constName string
	if s[1] == "(" && s[2] != "" {
		constName = s[2]
	}

	return utils.New[types.InterpretedCallFunc](utils.PropertiesAny{
		{
			Key:   "Name",
			Value: name,
		},
		{
			Key:   "Value",
			Value: getConstName(result, constName),
		},
	}, handleError), len(s)
}

func Create[F types.Func](t *types.InterpretedCallFunc) (lastInstruction *types.Instruction[F]) {
	newConst := types.NewConst[F](
		t.Value.Name,
		t.Value.Value,
	)

	if newConst != nil {
		lastInstruction = types.NewFunctionCall[F](t.Name, newConst)
	}

	return
}
