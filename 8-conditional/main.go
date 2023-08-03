package main

func main() {

	var age1 uint8 = 50

	if age1 >= 18 {
		println("eligible for vote")
	} else {
		println("not eligible for vote")
	}

	//  ticket in a bus
	// state = ap  age is <5 height is <100 cm then half ticket
	// state =ap full ticket
	// state = karnataka  age is <5 height is <120 cm then half ticket
	// state = karnataka gender == female then noticket
	// state = tamilnadue gender == female not ticket
	// stete = tamilnadu age <6 heigth is <120cm then half ticket or fullticket
	var (
		state  string
		gender rune // 'm' or 'f'
		age    uint8
		height float32
	)

	//state, gender, age, height = "tamilnadu", 'f', 4, 95.4
	state, gender, age, height = "tamilnadu", 'm', 7, 95.
	//println((state == "karnataka" || state == "tamilnadu") && gender == 'f')

	if (state == "karnataka" || state == "tamilnadu") && gender == 'f' {
		println("no ticket for ", state)
	} else if state == "karnataka" && age < 5 && height < 110 {
		println("half ticket for ", state)
	} else if state == "tamilnadu" && age < 6 && height < 120 {
		println("half ticket for ", state)
	} else if state == "ap" && age < 5 && height < 100 {
		println("halfticket for", state)
	} else {
		println("full ticket for ", state)
	}

}

// true && true --> true
// true && false -> false
// false && true -> false
// false && false -> false

// true || true -> true
// true || false -> true
// false || true -> true
// false || fasle -> fasle
