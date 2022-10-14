package utils

import (
	"errors"
	"fmt"
	"reflect"
)

type (
	MultiType any

	Property[T MultiType] struct {
		Key   string
		Value T
	}

	Properties[T MultiType] []Property[T]

	PropertiesAny []Property[any]
)

type ErrorHandler func(error)

func New[T any](props []Property[any], onError ErrorHandler) *T {
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
			fmt.Sprintf("New[T %T](props []Property[any]) *%T", t, t),
			false,
		)
	}

	return &t
}

func NewEmpty[T any]() *T {
	return new(T)
}
