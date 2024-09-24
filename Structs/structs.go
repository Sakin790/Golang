package main

import (
	"fmt"
)

type oder struct {
	id     int
	amount int
	status string
}

type student struct {
	id       int
	name     string
	location string
	class    int
	gender   string
}

func (o *student) changeGender(gender string) {
	o.gender = gender
}

func main() {

	myOder := oder{
		id:     1,
		amount: 5000,
		status: "recived",
	}

	fmt.Println(myOder)
	fmt.Println(myOder.id)

	mahid := student{
		id:       1,
		name:     "mahid",
		location: "Raj",
		class:    7,
		gender:   "Male",
	}
	mahid.changeGender("Female")
	fmt.Println(mahid)

	abir := student{
		id:       2,
		name:     "abir",
		location: "CTG",
		class:    9,
		gender:   "Male",
	}
	abir.changeGender("G")
	fmt.Println(abir)



	
	type person struct {
		firstname string
		lastname  string
		age       int
	}
	type address struct {
		Home     int
		Location string
		State    string
	}
	type contact struct {
		Email string
		Phone int
	}

	type Employee struct {
		Personal_Details person
		Emp_Address      address
		Emp_Contact      contact
	}

	var employee1 Employee
	employee1.Personal_Details = person{
		firstname: "Mahid",
		lastname:  "islam",
		age:       22,
	}
	employee1.Emp_Address = address{
		Home:     12,
		Location: "BD",
		State:    "Bangladesh",
	}
	employee1.Emp_Contact = contact{

		Email: "use1@gmail.com",
		Phone: 123212321,
	}
	fmt.Println(employee1)
}
