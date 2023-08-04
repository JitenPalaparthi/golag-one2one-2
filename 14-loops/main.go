package main

func main() {

	println("loop with multiple inits and conditions")
	for i, j := 1, 10; i <= 10 && j <= 20; i, j = i+1, j+2 {
		println("i: ", i, "j: ", j)
	}

	println("nested loop")
	count := 0
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			count++
			println("i: ", i, "j: ", j)
		}
	}

	println("total iterations:", count)

outer: // label any valid name(identifier) can be given
	for i := 1; i <= 10; i++ {
		for j := 3; j <= 10; j++ {
			if i == j {
				break outer
			}
			println("i: ", i, "j: ", j)
		}
	}

	// goto

	println("Goto--")

	i := 1
loop:
	if i%2 == 0 {
		print(i, " ")
	}
	i++
	if i <= 10 {
		goto loop
	}

	println("\nI am out of goto statement")

}
