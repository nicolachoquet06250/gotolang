package consts

import (
	"gotolang/types"
	"gotolang/utils"
)

func handleError(err error) {
	println(err.Error())
}

func Interpret(arr *[][]string, line, min, max int) (*types.InterpretedConst, int) {
	s := (*arr)[line][min:max]

	var name string

	if s[1] != "" {
		name = s[1]
	}

	var props = new(utils.PropertiesAny)

	if s[3] == `"` && s[4] != "" {
		props = &utils.PropertiesAny{
			{
				Key:   "Name",
				Value: name,
			},
			{
				Key:   "Type",
				Value: types.ICT_STRING,
			},
			{
				Key:   "Value",
				Value: s[4],
			},
		}
	} else if s[3] != `"` {
		props = &utils.PropertiesAny{
			{
				Key:   "Name",
				Value: name,
			},
			{
				Key:   "Type",
				Value: types.ICT_INT,
			},
			{
				Key:   "Value",
				Value: s[3],
			},
		}
	}

	if props != nil {
		return utils.New[types.InterpretedConst](*props, handleError), max
	}

	return nil, 0
}

func Create[F types.Func](t *types.InterpretedConst) (lastInstruction *types.Instruction[F]) {
	lastInstruction = types.NewConst[F](t.Name, t.Value)

	return
}
