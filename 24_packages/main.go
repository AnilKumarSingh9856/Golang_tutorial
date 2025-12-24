package main

import (
	"fmt"

	"github.com/project/auth"

	"github.com/project/users"
	
	"github.com/fatih/color"
)


func main() {
	
	auth.LoginWithUserCredentials("Anil", "12345")

	session := auth.GetSession()

	fmt.Println(session)

	user := users.User{
		Name: "Akash Singh",
		Email: "12345@gmail.com",
	}

	fmt.Println(user.Name, user.Email)

	color.Red(user.Email)
}