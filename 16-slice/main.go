package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 10)

	fmt.Println(slice1)

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}

	fmt.Println(slice1)
	fmt.Println("Sum of elements", SumOf(slice1))
	fmt.Println("Biggest Number:", BiggestNum(slice1))

	arr1 := [5]int{10, 11, 12, 13, 14}

	slice2 := arr1[:] // all elements an array is converted to a slice

	slice3 := arr1[:3] // 0th to the 3rd element

	slice4 := arr1[2:5] // 2nd to the 5th element

	slice5 := arr1[3:] // from 3rd element to the end

	slice6 := slice1[0:4]

	fmt.Println(slice2, slice3, slice4, slice5, slice6)
	slice7 := slice1 // This is a shallow copy ; Memory is copied
	fmt.Println(cap(slice1))
	fmt.Println(slice7)
	slice1 = append(slice1, 101)
	fmt.Println(slice7, slice1)
}

func SumOf(slice []int) int {
	sum := 0

	for _, v := range slice {
		sum += v
	}
	return sum
}

func BiggestNum(slice []int) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	big := slice[0]

	for i := 1; i < len(slice); i++ {
		if slice[i] > big {
			big = slice[i]
		}
	}
	return big
}
