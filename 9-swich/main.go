package main

func main() {

	var day int8 = 4

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	println("empty switch")

	var (
		state  string
		gender rune // 'm' or 'f'
		age    uint8
		height float32
	)

	//state, gender, age, height = "tamilnadu", 'f', 4, 95.4
	state, gender, age, height = "ap", 'f', 7, 95.

	switch {
	case (state == "karnataka" || state == "tamilnadu") && gender == 'f':
		println("no ticket for", state)
	case state == "karnataka" && age < 5 && height < 110:
		println("no ticket for", state)
	case state == "tamilnadu" && age < 6 && height < 120:
		println("no ticket for", state)
	case state == "ap" && age < 5 && height < 100:
		println("no ticket for", state)
	default:
		println("full ticket for", state)
	}

	char := '诶'

	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		println("The char is vovel")
	//case ''B','b':
	default:
		println("char is either a consonent or a special charcter")
	}

	// break is by default

}
