package main

import (
	"fmt"
	"reflect"
)

// in Go a type cannot be implicitely converted to another type.
// In order to do that , you need to do type casting

func main() {

	var num1 int

	var num2 uint8 = 127

	num1 = int(num2) // I want to store 1 byte in 8 bytes variable

	var num3 int = 1231 //

	var num4 uint8 = uint8(num3)

	fmt.Println(num1, num2, num3, num4)

	var float1 float32 = 101.1231           // float32
	var float2 = 12312312312.34312312321312 // float64

	float3 := float64(float1) + float2 // float64
	fmt.Println(float3)

	fmt.Printf("%.2f\n", float3)

	float4 := float64(float1 + float32(float2)) // not recommended still can do this
	fmt.Println(float4)
	fmt.Printf("%.2f\n", float4)
	//num1 = (int)num2
	// always convert the smaller size to bigger size so that the data does not stripe out

	var num5 uint8 = 65

	var str1 string = string(num5)

	// "67"

	fmt.Println(str1)

	str2 := 'A'

	num6 := int(str2)

	fmt.Println(num6)

	var num7 uint8 = 65

	var str3 string = fmt.Sprint(num7) // converts any data to string

	{
		var a, b int16
		var c int32
		a = 32767
		b = 32767
		//c = int64(a) + int64(b)
		c = int32(a) + int32(b)
		fmt.Println(c)

		//c = 131313
	}

	fmt.Println(str3, reflect.TypeOf(str3))
	fmt.Printf("value %s type %T\n", str3, str3)

	// str := "Hello world how are you doing"
	// str2:= "Hello folks"
	// str3:="Hello Gophers"

	// str = "fjfsdjf;dsjf'dl;fjld;sfjsfslfsdfjsfpoewjropwer"

	// int --> on 32 bit machines 4 bytes and 64 bit machines 8 bytes

}

// int a=100;
// int b=200;

// 1 2 4 8 16 32 64 128

//*int ptr = malloc(100) // to crete memory of 100bytes in heap

// free(*ptr)
