package main

import (
	_ "time" // if you dont use a package dont import it. Use _ (blank identifier if it is required to be presented)
)

var (
	i int = 1111
)

func main() {
	println(i)
	println("\nShadowing ------")
	// shadowing
	{ //-------------------------------------------------- i
		i := 10000
		println(i)
		{ //-------------------------------------- i
			i := 20000
			println(i)
		} //-------------------------------------- i
		println(i)
	} //-------------------------------------------------- i

}
