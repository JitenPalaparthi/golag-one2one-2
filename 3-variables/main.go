package main

import "fmt"

func main() {

	var num1 = 100

	//fmt.Println("Value of num1", num1)

	num2 := 101

	num3 := 12.45

	//age := 38// uint8

	var age uint8 = 38

	ok1 := false

	var ok2 bool

	str := "Hello World"

	rune1 := 'A'

	fmt.Println(num1, num2, num3, age, ok1, ok2, str, rune1)

	var (
		num4, num5, num6         = 101, 101.11, 12312312.23
		float1           float32 = 12.34
		float2           float64 = 14.56
	)

	fmt.Println(num4, num5, num6, float1, float2)

	const pi1 float32 = 3.14 // constants are evaluated at compile time

	const pi2 float32 = 3.0 + .14 // if the caluclation is on constants it works.

	//var float3 float32 = .14

	const float3 float32 = .14

	const pi3 float32 = 3.0 + float3 // it does not work. float3 is a variable that is not a const.

	const pi4 = 3.14

	const (
		MAXPRCS   = 4
		MAXTIME   = 10
		MAXDAYS   = 7
		MAXMONTHS = 12
	)
	//pi = 3.145

	fmt.Println(pi1, pi2, pi3, pi4, MAXDAYS, MAXMONTHS, MAXPRCS, MAXTIME)

}
