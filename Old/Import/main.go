package main

//packageName/foldername
//packageName/folde1/folder2/nested go on
import (
	"backend/cmd"
	"backend/cmd/database"
	"backend/internals"
)

func main() {
	cmd.PrintHelloName("sakin")
	internals.ShowDetails("mahid", 23)
	database.ConnectDB()
}
