package types

type InstructionValueType struct {
	InterpretedConstType
}

func CastToInstructionValueType(i InterpretedConstType) InstructionValueType {
	return InstructionValueType{InterpretedConstType: i}
}

func (t InstructionValueType) String() string {
	switch t.InterpretedConstType {
	case ICT_STRING, ICT_INT:
		return string(t.InterpretedConstType)
	default:
		return ""
	}
}

type InstructionType string

var (
	ASSIGN_CONST InstructionType = "assign_const"
	CALL_FUNC    InstructionType = "call_func"
)

func (it InstructionType) Is(instructionType InstructionType) bool {
	switch instructionType {
	case ASSIGN_CONST, CALL_FUNC:
		return true
	default:
		return false
	}
}

type Instruction[T comparable] struct {
	Name            string
	Action          Action
	Value           T
	ValueType       InstructionValueType
	Content         *Instruction[T]
	InstructionType InstructionType
}

func NewConst[T comparable](name string, action Action, value T, valueType InstructionValueType) *Instruction[T] {
	if valueType.IsValid() {
		return &Instruction[T]{
			Name:            name,
			Action:          action,
			Value:           value,
			Content:         nil,
			InstructionType: ASSIGN_CONST,
			ValueType:       valueType,
		}
	}
	return nil
}

func NewFunctionCall[T comparable](name string, content *Instruction[T]) *Instruction[T] {
	return &Instruction[T]{
		Name:            name,
		Content:         content,
		InstructionType: CALL_FUNC,
	}
}
