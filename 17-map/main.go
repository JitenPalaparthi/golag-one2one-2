package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var map1 map[string]int     //declarations
	map1 = make(map[string]int) // make is a built in func used to instantiate a map, slice and chan

	map1["bangalore-1"] = 560086
	map1["bangalore-2"] = 560096
	map1["bangalore-3"] = 560034
	// access map using range loop

	for key, value := range map1 {
		fmt.Println("Key:", key, "Value:", value)
	}

	val1, ok1 := map1["bangalore-1"]

	if ok1 {
		fmt.Println(val1)
	}

	val2, ok2 := map1["bangalore-5"]

	if ok2 {
		fmt.Println(val2)
	} else {
		fmt.Println("no value with the given key", val2)
	}

	slice1 := make([]int, 100)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100) // fill the slice with random numbers
	}
	// duplicate elements in the slice along with how many time those elements are
	fmt.Println(slice1)
	// To check the duplicate elements in a slice, the easiest way is by using map

	map2 := make(map[int]int)

	for _, v := range slice1 {
		v1, _ := map2[v]
		map2[v] = v1 + 1
	}
	fmt.Println(map2)

	delete(map1, "bangalore-1") // delete a key and associated value from the map

	fmt.Println(map1)

	map3 := make(map[string]any)

	map3["560086"] = "Bangalore-1,Yeshvantpur" // string
	map3["560086-isbangalore"] = true          // bool
	map3["560086-population"] = 3131232        // int

}

// O(1)
// maps are not ordered
// maps work on general hash function
// There are keys and values
// a hash key is applied on key
// maps are by default it contains 8 buckets
//   52b06cbe012db10887b9c659c89ff0b8  segment -> a mathematical number 4
//   a19414b055a3c668cb677f66576ab1c9  offset  -> a mathematical number 3
