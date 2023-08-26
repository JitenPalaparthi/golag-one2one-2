package main

import "fmt"

func main() {
	var e1 Employee // declaration

	fmt.Println("Normal struct example")
	e1.Id = 101 // give elements to the struct variable
	e1.Name = "Prabhaker"
	e1.Email = "Prabhakar.aluri@mail.com"
	e1.Address = map[string]string{"City": "Bangalore", "Country": "Bharat", "PinCode": "560096"} // creating a map with data and assigning to the field in the stuct

	e2 := Employee{Id: 102, Name: "Jiten", Email: "JItenp@outlook.com"} // declaring a struct and assigning elements to the struct

	fmt.Println(e1, e2)
	fmt.Println()
	fmt.Println("Composition of struct example")
	s1 := Student{Id: 101, Name: "Ravi", Email: "Ravi@gmail.com", Addr: Address{City: "Hyd", Country: "Bharat", PinCode: "543345"}}
	fmt.Println(s1)

	fmt.Println(s1.Name)
	fmt.Println(s1.Addr.City)
	fmt.Println(s1.Addr.Country)
	fmt.Println()
	fmt.Println("Promoted fields example")
	p1 := Person{Id: 101, Name: "Raju", Email: "Raju@Gmail.com", Address: Address{City: "Bangalore", Country: "Bharat", PinCode: "560086"}}

	fmt.Println(p1.Name)
	fmt.Println(p1.City)
	fmt.Println(p1.Country)
	fmt.Println()
	fmt.Println("Embedded struct fields example")

	c1 := Company{Id: 101, Name: "Aka-Labs", Email: "jp@aka-labs.io", Address: Address{City: "Guntur", Country: "Bharat", PinCode: "522316"}}

	fmt.Println(c1.Name)
	fmt.Println(c1.Address.City)

	fmt.Println()

	fmt.Println("Struct with anonymous fields")
	b1 := Branch{Id: 101, Name: "Aka-Labs", Email: "jp@aka-labs.io"}
	b1.Address = Address1{101, "Bangalore,Bharat"} // no field names here just field types
	fmt.Println(b1)

	fmt.Println()
	fmt.Println("Customised/anonymous inline struct")
	cm1 := struct {
		Id          int
		Name, Email string
		Address     struct {
			City, Country, PinCode string
		}
	}{
		Id:      101,
		Name:    "Jiten",
		Email:   "Jitenp@outlook.com",
		Address: Address{City: "Blr", Country: "Bharat", PinCode: "560086"},
	}

	fmt.Println(cm1)

}

type Employee struct { // Custom type or User defined type
	Id      int
	Name    string
	Email   string
	Address map[string]string // map as an address
}

type Student struct {
	Id    int
	Name  string
	Email string
	Addr  Address // Using another struct field
}

type Person struct {
	Id      int
	Name    string
	Email   string
	Address // promoted field. If promoted field is used then directly with object of Person can call all the fields and methods of promoted type
}
type Address struct {
	City, Country, PinCode string
}

type Company struct {
	Id          int
	Name, Email string
	Address     struct { // emedded struct
		City, Country, PinCode string
	}
}

type Branch struct {
	Id          int
	Name, Email string
	Address     Address1
}

type Address1 struct {
	int
	string
}
