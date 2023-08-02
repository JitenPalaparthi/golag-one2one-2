package main

import (
	"fmt"
	"reflect"
)

// type integer = int //typedef
//type any = interface{}

func main() {
	//Object obj= 123

	//var data interface{} // empty interface 1.18

	var data any

	data = 100

	//var num1 int = int(data) // type casting from any to int.. does not work

	var num1 int = data.(int) // type assertion var.(type)

	fmt.Println(num1)

	fmt.Println("Value:", data, "Type", reflect.TypeOf(data))

	data = true
	fmt.Println("Value:", data, "Type", reflect.TypeOf(data))

	data = 25.25
	fmt.Println("Value:", data, "Type", reflect.TypeOf(data))

	data = "Hello World!"
	fmt.Println("Value:", data, "Type", reflect.TypeOf(data))

	// var num1 integer = 100

	// fmt.Println(num1)

	var num2 int = 201

	var any1 any = num2

	var float1 float32 = 12.456

	any1 = float1

	float1 = any1.(float32)

	var num3 int = 100

	var float2 float32 = 12.43

	var any2 any = 1000 //int

	var any3 any = 12312.4343 // float64

	var float3 float64 = 12312313.1231231

	//var any4 any = false

	sum := float64(num3) + float64(float2) + float64(any2.(int)) + any3.(float64) + float3

	fmt.Println("Sum of:", sum)

	fmt.Printf("Sum of %.4f", sum)

}

// boxing and unboxing
