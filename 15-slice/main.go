package main

import "fmt"

func main() {

	var slice1 []int // declaration
	// fmt.Println(slice1)

	// if slice1 == nil {
	// 	fmt.Println("slice1 is a nil")
	// }

	slice1 = make([]int, 4) //
	// length is 4 and capacity is 4
	// type inference work for a slice
	slice1[0] = 100
	slice1[1] = 200
	slice1[2] = 300
	slice1[3] = 400
	//slice1[4] = 500 // this is not a legitimate statement.
	//Becase the lenght of the slice is only 4 but you are trying to assign value to the fifth element
	fmt.Println(slice1)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}
	fmt.Println("Length of a slice", len(slice1), "Capacity of a slice", cap(slice1))
	slice1 = append(slice1, 500)
	fmt.Println("Length of a slice", len(slice1), "Capacity of a slice", cap(slice1))
	slice1 = append(slice1, 600)
	slice1 = append(slice1, 700)
	fmt.Println("Length of a slice", len(slice1), "Capacity of a slice", cap(slice1))
	slice1 = append(slice1, 800)
	fmt.Println("Add of first element of slice:", &slice1[0])
	slice1 = append(slice1, 900)
	fmt.Println("Length of a slice", len(slice1), "Capacity of a slice", cap(slice1))
	fmt.Println("Add of first element of slice:", &slice1[0])
}

// Slice is a growable array. It can grow or can shrink
// make is a built in function that is used to allocate slice, map and channel
// slice has length and capacity; Unless given initially length and capacity are same
// how to append an element to the slice
