package main

import (
	"fmt"
	"gotolang/parser"
	"gotolang/syntax_interpreters"
	"gotolang/types"
	"gotolang/utils"
	"os"
	"reflect"
)

type (
	MultiType interface{}

	AssocList[T MultiType] map[string]T

	Toto struct {
		Name string
		Age  int
	}

	Tata struct {
		FullName string
		Notes    []float32
	}
)

func toto[T MultiType](data AssocList[MultiType]) (t *T) {
	//var t = new(T)

	for _, val := range data {
		for k, v := range val.(AssocList[MultiType]) {
			f := reflect.ValueOf(v)
			st := reflect.ValueOf(t)

			if st.Kind() == reflect.Struct {
				s := st.FieldByName(k)

				if f.Kind() == reflect.String && true == s.CanSet() {
					s.SetString(f.String())
				} else if f.Kind() == reflect.Float64 && true == s.CanSet() {
					s.SetFloat(f.Float())
				} else if f.Kind() == reflect.Slice && true == s.CanSet() {
					s.Set(f.Slice(0, f.Len()))
				}
			}
		}
	}

	return
}

func tata() {
	var p AssocList[MultiType]

	println(fmt.Sprintf("%T", toto[Toto](p)))
	println(fmt.Sprintf("%T", toto[Tata](p)))
}

func main() {
	file := os.Args[1]
	data := utils.OpenFile(file)

	var splitCode = parser.Parse(data)

	utils.DebugAction(func() {
		for _, row := range *splitCode {
			for _, col := range row {
				println(col)
			}
		}
	}, "", false)

	syntax_interpreters.Interpret(
		types.NewProgram(splitCode),
	)

	utils.DebugAction(func() {
		tata()
	}, "Test de types génériques avec remplissage dynamique", false)
}
