package utils

import (
	"errors"
	"fmt"
	"reflect"
)

type (
	MultiType interface{ interface{} | []interface{} }

	Property[T MultiType] struct {
		Key   string
		Value T
	}
)

type ErrorHandler func(error)

func GetFromType[T MultiType](props []Property[any], onError ErrorHandler) *T {
	var t = *new(T)
	var newT = reflect.ValueOf(&t).Elem()

	defer func() {
		if err := recover(); err != nil {
			onError(errors.New(fmt.Sprintln(err)))
		}
	}()

	for _, property := range props {
		newT.FieldByName(property.Key).Set(reflect.ValueOf(property.Value))

		DebugAction(
			func() {
				println(fmt.Sprintf("%s => %T", property.Key, property.Value))
			},
			fmt.Sprintf("GetFromType[T %T](props []Property[any]) *%T", t, t),
			false,
		)
	}

	return &t
}
