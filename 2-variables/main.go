package main

import "fmt"

var num int = 101 // global variable

func main() {
	var num1 int // local to entire main func
	{
		var num2, num3 float32 = 12.34, 12.45 // local to the current scope
		fmt.Println("Value of num2", num2, num3)
	}

	//num2 = 120.34

	var str string = "Hello World"

	println("value of num1", num1)
	println(str)

	num1 = 100

	println("value of num1", num1)

}

// numbers
// int,int8,int16,int32,int64 , uint8,uint16,uint32,uint64  -> default value is 0
// float32,float64 > default value is 0
// rune alias for int32 > default value is 0
// byte alias for int8 > default value is 0
// bool > default value is false
// string > default value is ""
// complex64, complex128
// uintptr
// interface{} alias is any

// type inference
// Based on a type of a variable a default value is given
