package main

import (
	"errors"
	"fmt"
)

// nil can be checked on slices, maps, channels, pointers , any and also on interfaces
func main() {

	var any1 any = false

	switch any1.(type) {
	case int:
		fmt.Println(any1, "is integer")
	case uint:
		fmt.Println(any1, "is usigied integer")
	case uint8:
		fmt.Println(any1, "is  unit8 type")
	case float32:
		fmt.Println(any1, "is float32 type")
	case string:
		fmt.Println(any1, "is of type string")
	case bool:
		fmt.Println(any1, "is of type bool")
	default:
		fmt.Println(any1, "is some other type")
	}

	sq, err := square(12.12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sq)
	}

	sq, err = square("Hello World")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sq)
	}

	sq, err = square(true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sq)
	}

}

func square(a any) (any, error) {
	var sq any
	switch a.(type) {
	case int:
		sq = a.(int) * a.(int)
	case uint:
		sq = a.(uint) * a.(uint)
	case float32:
		sq = a.(float32) * a.(float32)
	case float64:
		sq = a.(float64) * a.(float64)
	case string:
		return sq, errors.New("for string cannot implement square")
	default:
		return sq, errors.New("cannot implement square on other types")
	}
	return sq, nil
}
