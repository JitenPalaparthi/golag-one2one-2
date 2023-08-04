package main

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}

	println("\neven")

	for i := 2; i <= 10; i += 2 {
		print(i, " ")
	}
	println("\neven")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}
	println("\neven")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue // continue is to continue with the next iteration
		}
		print(i, " ")
	}
	println("\neven")

	i := 1

	for ; i <= 10; i++ {
		if i%2 == 1 {
			continue // continue is to continue with the next iteration
		}
		print(i, " ")
	}

	println("\n even")
	i = 0 /// initialiser

	for {
		i++
		if i > 10 {
			break // break breaks the entire loop
		}

		if i%2 == 1 {
			continue
		}

		print(i, " ")
		// incrementor
	}

}

// go has only for loop
