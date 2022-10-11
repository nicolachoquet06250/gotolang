package types

type InterpretedCreatedFuncType string

/*var (
	ICRFT_STRING InterpretedConstType = "string"
	ICRFT_INT    InterpretedConstType = "int"
)*/

func (ict InterpretedCreatedFuncType) IsValid() bool {
	switch ict {
	/*case ICT_STRING, ICT_INT:
	return true*/
	default:
		return true
	}
}

type ParametersType map[string]string

type InterpretedCreatedFunc struct {
	Name       string
	Parameters ParametersType
	Content    string
	ReturnType string
}
