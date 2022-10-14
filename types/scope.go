package types

import (
	"errors"
	"gotolang/utils"
)

type ScopeType string

var (
	FUNCTION  ScopeType = "func"
	ANONYMOUS ScopeType = "noname"
	MAIN      ScopeType = "main"
)

func (st ScopeType) IsValid() bool {
	switch st {
	case FUNCTION, ANONYMOUS, MAIN:
		return true
	default:
		return false
	}
}

type Scope[F Func] struct {
	Name         string
	Type         ScopeType
	Instructions []Instruction[F]
}

func (s *Scope[F]) AddInstruction(instruction Instruction[F]) *Scope[F] {
	s.Instructions = append(s.Instructions, instruction)
	return s
}

func NewScope[F Func](name string, scopeType ScopeType) (scope *Scope[F], err error) {
	if scopeType.IsValid() {
		scope = &Scope[F]{
			Name:         name,
			Type:         scopeType,
			Instructions: []Instruction[F]{},
		}
	} else {
		err = errors.New("invalid scope type")
	}
	return
}

type Program[F Func] struct {
	Scope[F]
	ParsedCode *[][]string
}

func NewProgram[F Func](parsedCode *[][]string) *Program[F] {
	var err error
	var programScope *Scope[F]
	programScope, err = NewScope[F]("main", MAIN)
	utils.CheckError(err)

	return utils.New[Program[F]](utils.PropertiesAny{
		{
			Key:   "Scope",
			Value: *programScope,
		},
		{
			Key:   "ParsedCode",
			Value: parsedCode,
		},
	}, func(err error) {
		println(err.Error())
	})

}
