package main

import "fmt"

func main() {
	Greet()
	Greet()
	GreetBy("Hello folks")
	GreetBy("Hello folks")

	var num1 int = 100

	Incr(num1) // call or pass by value

	fmt.Println(num1)

	a1, b1 := 10, 20

	c1 := Sum(a1, b1)

	fmt.Println(c1)

	sum1, sub1, mul1, div1, mod1 := Calc(a1, b1)

	fmt.Println(sum1, sub1, mul1, div1, mod1)

	sum2, _, mul2, div2, _ := Calc(201, 301) // blank identifier

	fmt.Println(sum2, mul2, div2)

	_, _, _, _, mod3 := Calc(200, 300)
	fmt.Println(mod3)

}

// func and methods
// methods are associated with types

func Greet() {
	println("Hello World")
}

func GreetBy(msg string) {
	println(msg)
}

func Incr(num int) { // creates a new variable and assigns a value to the new variable inside Function stack frame
	num++
}

func Sum(a, b int) int {
	c := a + b
	return c
	// d := c + c
	// fmt.Println(d)
	// return d
}

// func Calc(a, b int) (int, int, int, float32, int) {
// 	return a + b, a - b, a * b, float32(a) / float32(b), a % b
// }

// func Calc1(a, b int) (int, int, int, float32, int) {
// 	sum := a + b // creating and assigning a value
// 	sub := a - b
// 	mul, div, mod := a*b, float32(a)/float32(b), a%b

// 	return sum, sub, mul, div, mod
// }

func Calc(a, b int) (sum int, sub int, mul int, div float32, mod int) {
	//return a + b, a - b, a * b, float32(a) / float32(b), a % b
	sum = a + b // sum variable is alaredy created so you are just assigning a value to it
	sub = a - b
	mul, div, mod = a*b, float32(a)/float32(b), a%b

	return sum, sub, mul, div, mod
}
