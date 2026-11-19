package main

import (
	"github.com/sakin790/myapp/auth"
)

func main() {
	auth.LoginWithCradential("Sakin", "password")
	auth.ShowStudentDetails()
	auth.UserSeason()

}
