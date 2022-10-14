package main

import (
	"fmt"
	. "gotolang/utils"
)

type (
	Toto struct {
		Name string
		Age  int
	}

	Tata struct {
		FullName string
		Notes    []float64
	}
)

func getErrorHandler(prefix string) ErrorHandler {
	return func(err error) {
		println(prefix, err.Error())
	}
}

func _main() {
	DebugAction(
		func() {
			t1 := New[Toto](
				[]Property[any]{
					{
						Key:   "Name",
						Value: "Nicolas",
					},
					{
						Key:   "Age",
						Value: 27,
					},
				},
				getErrorHandler("Erreuuuuuuuur !!!! "),
			)

			if t1 != nil {
				println(fmt.Sprintf("%T", t1))
			}

			t2 := New[Tata](
				[]Property[any]{
					{
						Key:   "FullName",
						Value: "Nicolas Choquet",
					},
					{
						Key:   "Notes",
						Value: []float64{12, 12.5, 20, 15},
					},
				},
				getErrorHandler("Nonnnnnnnnnnnnnnn !!!! "),
			)

			if t2 != nil {
				println(fmt.Sprintf("%T", t2), "\n_______________________________________________")
			}

			t3 := New[Tata](
				[]Property[any]{
					{
						Key:   "FullName",
						Value: "Nicolas Choquet",
					},
					{
						Key:   "Notes",
						Value: []float64{12, 12.5, 20, 15},
					},
				},
				getErrorHandler("Ouiiiiiiiiiiiiiiiiii !!!! "),
			)

			if t3 != nil {
				println("FullName: " + t3.FullName)
				println("Notes length: " + fmt.Sprintf("%v", len(t3.Notes)))
			}
		},
		"Test de types génériques avec remplissage dynamique",
		true,
	)
}
