package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	typeOf := reflect.TypeOf(x)
	fmt.Println("Type:", typeOf)
	// Type: float64
	v := reflect.ValueOf(x)
	fmt.Println("Value:", v)
	// Value: 3.4
	fmt.Println("Type of value:", v.Type())
	// Type of value: float64

	// 使用指针
	vp := reflect.ValueOf(&x)
	// 获取指针指向的值
	v = vp.Elem()
	v.SetFloat(7.1)
	fmt.Println(x)

}
