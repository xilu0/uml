package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func printFields(person interface{}) {
	val := reflect.ValueOf(person)

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}

func setAge(person interface{}, newAge int) {
	val := reflect.ValueOf(person).Elem()
	ageField := val.FieldByName("Age")

	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(int64(newAge))
	}
}

func main() {
	p := Person{Name: "John Doe", Age: 30}
	fmt.Println("Before setting age:")
	printFields(p)

	setAge(&p, 31) // 注意：传递p的指针
	fmt.Println("After setting age:")
	printFields(p)
}
