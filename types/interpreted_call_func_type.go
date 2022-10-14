package types

type InterpretedCallFuncType string

var (
	ICFT_STRING InterpretedCallFuncType = "string"
	ICFT_INT    InterpretedCallFuncType = "int"
	ICFT_CONST                          = InterpretedCallFuncType(CONST)
)

func (icft InterpretedCallFuncType) IsValid() bool {
	switch icft {
	case ICFT_STRING, ICFT_INT, ICFT_CONST:
		return true
	default:
		return false
	}
}

type InterpretedCallFunc struct {
	Name  string
	Value *InterpretedConst
}
