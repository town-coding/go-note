### GO 使用reflect
反射（Reflection）是一种在运行时检查类型、结构和变量值的能力
#### 反射的核心类型
Go 的反射机制主要依赖以下两个核心类型：

1. reflect.Type：用于表示 Go 类型的元数据，可以通过 reflect.TypeOf() 获取。
2. reflect.Value：表示存储在某个变量中的实际值，通过 reflect.ValueOf() 获取。

#### 基本用法
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 获取类型信息
	// 可以通过 reflect.TypeOf() 获取变量的类型信息
	var x float64 = 3.4
	typeOf := reflect.TypeOf(x)
	fmt.Println("Type:", typeOf)
	// Type: float64
	
	// 获取值信息
	//可以通过 reflect.ValueOf() 获取变量的值：
	v := reflect.ValueOf(x)
	fmt.Println("Value:", v)
	// Value: 3.4
	fmt.Println("Type of value:", v.Type())
	// Type of value: float64

	// 使用指针
	vp := reflect.ValueOf(&x)
	// 获取指针指向的值
	v = vp.Elem()
	// 修改目标值，只有通过指针获取的 reflect.Value 才能修改原始值
	v.SetFloat(7.1)
	fmt.Println(x)
}
```

#### 反射中处理结构体
```go
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
	// 使用 reflect.Type 和 reflect.Value 可以获取结构体的字段信息
	p := Person{Name: "Alice", Age: 25}
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field: %s, Type: %s, Value: %v\n", field.Name, field.Type, value.Interface())
	}
	// Field: Name, Type: string, Value: Alice 
	// Field: Age, Type: int, Value: 25
	
	vp := reflect.ValueOf(&p).Elem()

	nameField := vp.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Bob")
	}

	if vp.FieldByName("Age").CanSet() {
		vp.FieldByName("Age").SetInt(10)
	}

	fmt.Println(p) // 输出：{Bob 25}
}
```
#### 反射的实际应用场景
反射在以下场景中非常有用：
1. JSON/XML 编解码：解析未知的结构体。
2. ORM 框架：动态生成 SQL 语句，处理数据库映射。
3. 编写通用库：例如日志库、序列化库，可以处理任意类型的数据。
#### 反射的性能开销
虽然反射非常强大，但由于反射是在运行时进行的，因此会带来一些性能开销。一般来说，反射的操作速度会比普通的类型操作慢。因此，在性能要求较高的场景下，应尽量避免频繁使用反射。

