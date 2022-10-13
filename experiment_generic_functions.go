package main

import (
	"encoding/json"
	"fmt"
	"gotolang/utils"
	"reflect"
)

type (
	MultiType interface{ interface{} | []interface{} }

	Property[T MultiType] struct {
		Key   string
		Value T
	}

	Toto struct {
		Name string
		Age  int
	}

	Tata struct {
		FullName string
		Notes    []float64
	}
)

func GetFromType[T MultiType](data []Property[any]) T {
	var t = *new(T)
	rv := reflect.ValueOf(&t).Elem()

	for _, property := range data {
		m, _ := json.Marshal(t)
		var x map[string]interface{}
		_ = json.Unmarshal(m, &x)

		switch fmt.Sprintf("%T", property.Value) {
		case "string":
			rv.FieldByName(property.Key).SetString(property.Value.(string))
			break
		case "bool":
			rv.FieldByName(property.Key).SetBool(property.Value.(bool))
			break
		case "int":
			println(fmt.Sprintf("%T", property.Value))
			rv.FieldByName(property.Key).SetInt(property.Value.(int64))
		case "int64":
			rv.FieldByName(property.Key).SetInt(property.Value.(int64))
			break
		case "uint64":
			rv.FieldByName(property.Key).SetUint(property.Value.(uint64))
			break
		case "float64":
			rv.FieldByName(property.Key).SetFloat(property.Value.(float64))
			break
		default:
			reflect.Copy(rv.FieldByName(property.Key), reflect.ValueOf(property.Value))
			break
		}
	}

	return t
}

func main() {
	utils.DebugAction(
		func() {
			println(fmt.Sprintf("%T",
				GetFromType[Toto](
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
				)),
				"\n_______________________________________________",
			)

			println(fmt.Sprintf("%T",
				GetFromType[Tata](
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
				)),
				"\n_______________________________________________",
			)

			/*t := GetFromType[Tata](
				[]Property[any]{
					{
						Key:   "FullName",
						Value: "Nicolas Choquet",
					},
					{
						Key:   "Notes",
						Value: []float32{12, 12.5, 20, 15},
					},
				},
			)
			println(fmt.Sprintf("%v", t))
			println(t.FullName, len(t.Notes))*/
		},
		"Test de types génériques avec remplissage dynamique",
		true,
	)
}
