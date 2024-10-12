package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	v := reflect.ValueOf(&p).Elem()
	tp := reflect.TypeOf(&p).Elem()
	// v 的数据类型
	println(tp.Kind())

	nameField := v.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Bob")
	}

	if v.FieldByName("Age").CanSet() {
		v.FieldByName("Age").SetInt(10)
	}

	fmt.Println(p) // 输出：{Bob 25}
}
