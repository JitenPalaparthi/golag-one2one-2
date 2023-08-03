package main

//fallthrough

var (
	number int
)

func main() {
	number = 6
	switch {
	case number%8 == 0:
		println(number, "is divisible by 8")
		fallthrough
	case number%4 == 0:
		println(number, "is divisible by 4")
		fallthrough
	case number%2 == 0:
		println(number, "is divisible by 2")
		fallthrough
	case number%3 == 0:
		println(number, " is divisible by 3")
	}
}
