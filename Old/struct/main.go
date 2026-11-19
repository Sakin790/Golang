package main

import "fmt"

type Addres struct {
	Country     string
	City        string
	Zip         int
	GeoLocation string
}

type User struct {
	Id     int
	Name   string
	Gender string
	Addres Addres
}

type Prople interface {
	ChangeZipCode(code int)
}

// no need to pass *, bcz ami kono value change korchi na
func (usr User) printUserName() {
	fmt.Println(usr.Name)
}

func (usr *User) chanageName(name string) {
	usr.Name = name
}

func (usr *User) ChangeZipCode(codes int) {
	usr.Addres.Zip = codes
}

func main() {

	mahid := User{
		Id:     1,
		Name:   "mahid",
		Gender: "male",
		Addres: Addres{
			Country:     "Bangladesh",
			City:        "Rajshahi",
			Zip:         6271,
			GeoLocation: "xyz",
		},
	}
	fmt.Println(mahid)
	fmt.Println(mahid.Addres.Zip)
	mahid.printUserName()
	mahid.chanageName("Sakin")
	mahid.printUserName()
}
