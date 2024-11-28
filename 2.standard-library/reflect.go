package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// type Sample struct {
// 	Name string `required:"true" max:"30"`
// }

// type Person struct {
// 	Name    string `required:"true" max:"30"`
// 	Address string `required:"true" max:"30"`
// 	Email   string `required:"true" max:"30"`
// }

// func ReadField(value any) {
// 	var valueType reflect.Type = reflect.TypeOf(value)
// 	fmt.Println("Type Name", valueType.Name())
// 	for i := 0; i < valueType.NumField(); i++ {
// 		var structField reflect.StructField = valueType.Field(i)
// 		fmt.Println(structField.Name, "with type", structField.Type)
// 		fmt.Println(structField.Tag.Get("required"))
// 		fmt.Println(structField.Tag.Get("max"))
// 	}
// }

// func IsValid(value any) (result bool) {
// 	result = true
// 	t := reflect.TypeOf(value)

// 	for i := 0; i < t.NumField(); i++ {
// 		f := t.Field(i)
// 		if f.Tag.Get("required") == "true" {
// 			data := reflect.ValueOf(value).Field(i).Interface()
// 			result = data != ""

// 			if !result {
// 				return result
// 			}
// 		}
// 	}

// 	return result
// }

// func main() {
// 	ReadField(Sample{"ilham"})
// 	ReadField(Person{"ilham", "indonesian", "email"})

// 	fmt.Println(IsValid(Person{
// 		Name:    "ilham",
// 		Address: "jln",
// 		Email:   "hello",
// 	}))
// }

type Account struct {
	Name     string `required:"true" min:"5"`
	Password string `required:"true" min:"8" has_number:"true"`
}

func IsValid(value any) bool {
	result := true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		data := reflect.ValueOf(value).Field(i).Interface()
		if f.Tag.Get("required") == "true" {
			result = data != ""
			if !result {
				return result
			}
		}
		min, err := strconv.Atoi(f.Tag.Get("min"))
		if err == nil && min > 0 {
			if str, ok := data.(string); ok {
				result = len(str) >= min
			}
			if !result {
				return result
			}
		}
		if f.Tag.Get("has_number") == "true" {
			has_number := false
			if str, ok := data.(string); ok {
				for _, s := range str {
					_, err := strconv.Atoi(string(s))
					if err == nil {
						has_number = true
					}
					if has_number {
						break
					}
				}
			}

			if result = has_number; !result {
				return result
			}
		}
	}

	return result
}

func main() {
	fmt.Println(IsValid(Account{
		Name:     "ilhamganteng",
		Password: "hellossss2hs",
	}))
}
