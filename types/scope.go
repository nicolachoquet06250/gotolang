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

type Scope struct {
	Name         string
	Type         ScopeType
	Instructions []Instruction[string]
}

func (s *Scope) AddInstruction(instruction Instruction[string]) *Scope {
	s.Instructions = append(s.Instructions, instruction)
	return s
}

func NewScope(name string, scopeType ScopeType) (scope *Scope, err error) {
	if scopeType.IsValid() {
		scope = &Scope{
			Name:         name,
			Type:         scopeType,
			Instructions: []Instruction[string]{},
		}
	} else {
		err = errors.New("invalid scope type")
	}
	return
}

type Program struct {
	Scope
	ParsedCode *[][]string
}

func NewProgram(parsedCode *[][]string) *Program {
	var err error
	var programScope *Scope
	programScope, err = NewScope("main", MAIN)
	utils.CheckError(err)

	return &Program{
		Scope:      *programScope,
		ParsedCode: parsedCode,
	}
}
