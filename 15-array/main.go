package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {

	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(100)
	}

	fmt.Println(arr1)

	var arr2 [2]int
	arr2[0] = 1
	arr2[1] = 2

	fmt.Println(arr2)

	arr3 := [3]int{10, 11, 12} // shorthand declaration of array
	fmt.Println(arr3)

	arr4 := [...]int{10, 11, 12} // shorthand declaration of array
	fmt.Println(arr4)

	// for range loop

	// for range loop can be used for array, slice , map , channel
	// for array and slice -> the first return type is index and second one is value

	for i, v := range arr4 {
		fmt.Println("Index", i, "Value:", v)
	}

	for _, v := range arr4 {
		fmt.Println("Value:", v)
	}

	for i, _ := range arr4 {
		fmt.Println("Index", i)
	}

	str1 := "Hello World!" //0-255 uint8

	for i := 0; i < len(str1); i++ {
		fmt.Print(string(str1[i]), " ")
	}
	fmt.Println()
	str2 := "你好罗尔德"
	//str2 := "Hello World"

	for i := 0; i < len(str2); i++ {
		fmt.Println("type of", reflect.TypeOf(str2[i]), string(str2[i]), " ")
	}

	// if a normal loop is used on strings, it iterates each byte
	// the problem arises when each char is more than byte
	fmt.Println()
	for _, v := range str2 {
		fmt.Println("type of ", reflect.TypeOf(v), string(v), " ")
	}

}

// arrays are collection of elements of the same type
// the size of the array is fixed and must to knwon at the compile time
// once an array is created you cannot grow or shrink the array
