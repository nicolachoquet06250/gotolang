package types

import (
	"gotolang/utils"
)

func handleError(err error) {
	println(err.Error())
}

type InstructionType string

const (
	CreateConst    InstructionType = "create_const"
	CreateFunction InstructionType = "create_function"
	CallFunction   InstructionType = "call_function"
)

func (it InstructionType) Is(instructionType InstructionType) bool {
	switch instructionType {
	case it:
		return true
	default:
		return false
	}
}

//------------------------------------------------------------------------------------------------

type (
	Func interface {
		func() | func(any) | func(...any) | func() any | func(any) any | func(...any) any | *any
	}
	TConst struct {
		Name  string
		Value any
	}
	TFunction[F Func] struct {
		Name       string
		ReturnType string
		Content    []Instruction[F]
		Ref        F
		Parameters []TConst
	}
	TCall[F Func] struct {
		Func       TFunction[F]
		Parameters []Instruction[F]
	}
	Instruction[F Func] struct {
		Type     InstructionType
		Const    *TConst
		Function *TFunction[F]
		Call     *TCall[F]
	}
)

func (i *Instruction[F]) IsCreateConst() bool {
	return i.Type == CreateConst
}
func (i *Instruction[F]) IsCreateFunction() bool {
	return i.Type == CreateFunction
}
func (i *Instruction[F]) IsCallFunction() bool {
	return i.Type == CallFunction
}

func NewConst[F Func, T any](name string, value T) *Instruction[F] {
	return utils.New[Instruction[F]](utils.PropertiesAny{
		{
			Key:   "Type",
			Value: CreateConst,
		},
		{
			Key: "Const",
			Value: utils.New[TConst](utils.PropertiesAny{
				{
					Key:   "Name",
					Value: name,
				},
				{
					Key:   "Value",
					Value: value,
				},
			}, handleError),
		},
	}, handleError)
}

func NewFunctionCall[F Func](name string, content *Instruction[F]) *Instruction[F] {
	return utils.New[Instruction[F]](utils.PropertiesAny{
		{
			Key:   "Type",
			Value: CallFunction,
		},
		{
			Key: "Call",
			Value: utils.New[TCall[F]](utils.PropertiesAny{
				{
					Key: "Func",
					Value: utils.New[TFunction[F]](utils.PropertiesAny{
						{
							Key:   "Name",
							Value: name,
						},
						{
							Key:   "ReturnType",
							Value: "null",
						},
						{
							Key:   "Content",
							Value: []Instruction[F]{},
						},
						{
							Key:   "Parameters",
							Value: []TConst{},
						},
					}, handleError),
				},
				{
					Key:   "Parameters",
					Value: []Instruction[F]{*content},
				},
			}, handleError),
		},
	}, handleError)
}

func NewCreatedFunction[F Func](name string, returnType string, parameters []TConst) *Instruction[F] {
	return utils.New[Instruction[F]](utils.PropertiesAny{
		{
			Key:   "Type",
			Value: CreateFunction,
		},
		{
			Key: "Function",
			Value: utils.New[TFunction[F]](utils.PropertiesAny{
				{
					Key:   "Name",
					Value: name,
				},
				{
					Key:   "ReturnType",
					Value: returnType,
				},
				{
					Key:   "Content",
					Value: utils.NewEmpty[[]Instruction[F]](),
				},
				{
					Key:   "Parameters",
					Value: parameters,
				},
			}, handleError),
		},
	}, handleError)
}

//@TODO Utiliser et corriger si besoin la nouvelle structure d'instructions
