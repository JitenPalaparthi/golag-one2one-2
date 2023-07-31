package main

import "fmt"

func main() {
	a1, b1 := 10, 20 // shorthand declaration of multiple variables

	t1 := b1
	b1 = a1
	a1 = t1

	a1, b1 = b1, a1 // swapping is easy in Go

	fmt.Println(a1, b1) // 10,20

	a2, b2, c2 := 10, 20, 30

	fmt.Println("Before swapping a2,b2,and c2", a2, b2, c2) // 10,20,30

	a2, b2, c2 = b2, c2, a2

	fmt.Println("After swapping a2,b2,and c2", a2, b2, c2) // 20,30,10

}
