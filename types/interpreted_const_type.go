package types

type InterpretedConstType string

var (
	ICT_STRING InterpretedConstType = "string"
	ICT_INT    InterpretedConstType = "int"
)

func (ict InterpretedConstType) IsValid() bool {
	switch ict {
	case ICT_STRING, ICT_INT:
		return true
	default:
		return false
	}
}

type InterpretedConst[T comparable] struct {
	Name  string
	Type  InterpretedConstType
	Value T
}
