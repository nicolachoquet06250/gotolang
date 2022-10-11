package create_function

import (
	"gotolang/types"
	"strings"
)

func Interpret(arr *[][]string, line, min, max int) (*types.InterpretedCreatedFunc, int) {
	s := (*arr)[line][min:max]

	params := make(types.ParametersType)
	name := ""

	if s[1] != "" {
		name = s[1]
	}

	content := ""
	returnType := "null"

	if strings.Contains(strings.Join(s, ""), "=>") {
		returnType = strings.Trim(
			strings.Split(
				strings.Replace(
					strings.Join(s, " "),
					";",
					"",
					1,
				),
				" = > ",
			)[1],
			" ",
		)
		content = returnType
	}

	started := false
	_params := ""
	for _, e := range s {
		if e == "(" && !started {
			started = true
			continue
		}
		if e != ")" && started {
			_params += e + " "
			continue
		}
		if e == ")" {
			started = false
			continue
		}
	}

	aParams := strings.Split(_params, ",")

	for _, e := range aParams {
		splitParam := strings.Split(strings.Trim(e, " "), ":")

		if len(splitParam) > 1 {
			params[splitParam[0]] = "any"
		} else {
			params[splitParam[0]] = splitParam[1]
		}
	}

	return &(types.InterpretedCreatedFunc{
		Name:       name,
		Parameters: params,
		Content:    content,
		ReturnType: returnType,
	}), max
}

func Create(t *types.InterpretedCreatedFunc) (lastInstruction *types.Instruction[string]) {
	println(t.Name, t.ReturnType, t.Content, len(t.Parameters))
	for _key, _type := range t.Parameters {
		println("Le paramètre " + _key + " est de type " + _type)
	}
	return
}
