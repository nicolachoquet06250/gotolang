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
		testGetFromType()
	}, "Test de types génériques avec remplissage dynamique", true)
}

type (
	MultiType interface{ any | []any }

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

func GetFromType[T MultiType](data AssocList[MultiType]) (t T) {
	for key, val := range data {
		if utils.MatchRegex(`^((\*?\[\])?[a-z0-9]+)$`, reflect.TypeOf(val).String()) {
			f := reflect.ValueOf(val)
			st := reflect.New(reflect.TypeOf(t))

			if st.Kind() == reflect.Struct {
				s := st.FieldByName(key)

				if f.Kind() == reflect.String && true == s.CanSet() {
					s.SetString(f.String())
				} else if f.Kind() == reflect.Float64 && true == s.CanSet() {
					s.SetFloat(f.Float())
				} else if f.Kind() == reflect.Slice && true == s.CanSet() {
					s.Set(f.Slice(0, f.Len()))
				}
			}

			var _t = st.Elem().Interface().(T)
			t = _t
		}
	}

	return
}

func testGetFromType() {
	println(
		fmt.Sprintf(
			"%T",
			GetFromType[Toto](
				AssocList[MultiType]{
					"Name": "Nicolas",
					"Age":  27,
				},
			),
		),
	)

	println(
		fmt.Sprintf(
			"%T",
			GetFromType[Tata](
				AssocList[MultiType]{
					"FullName": "Nicolas Choquet",
					"Notes":    []float32{12, 12.5, 20, 15},
				},
			),
		),
	)

	t := GetFromType[Tata](
		AssocList[MultiType]{
			"FullName": "Nicolas Choquet",
			"Notes":    []float32{12, 12.5, 20, 15},
		},
	)
	println(fmt.Sprintf("%v", t))
	println(t.FullName, len(t.Notes))
}
