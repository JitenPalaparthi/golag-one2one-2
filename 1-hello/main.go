package main

import "fmt"
import "os"
import "time"

func main() {
	println("Hello World!")
	fmt.Println("Hello World!", time.Now())
	os.Exit(0) // successful exit
}

// go runtime can recognise main func as the starting point and that main func must be in main pacakge.
