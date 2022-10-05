package consts

type InterpretedConstType string

var (
	STRING InterpretedConstType = "string"
	INT    InterpretedConstType = "int"
)

type InterpretedConst[T comparable] struct {
	Name  string
	Type  InterpretedConstType
	Value T
}

func Interpret(arr *[][]string, line, min, max int) *InterpretedConst[string] {
	s := (*arr)[line][min:max]

	var name string

	if s[1] != "" {
		name = s[1]
	}

	if s[3] == `"` && s[4] != "" {
		var v = s[4]
		return &(InterpretedConst[string]{
			Name:  name,
			Type:  STRING,
			Value: v,
		})
	} else if s[3] != `"` {
		var v = s[3]
		return &(InterpretedConst[string]{
			Name:  name,
			Type:  INT,
			Value: v,
		})
	}

	return nil
}
