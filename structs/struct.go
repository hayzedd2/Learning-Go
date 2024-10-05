package main

import (
	"fmt"

	"example.com/investmentCalculator/structs/userhooks"
	"github.com/Pallinder/go-randomdata"
)


func main() {
	firstName := getUserData("Enter your firstname:")
	lastName := getUserData("Enter your lastname: ")
	birthDate := getUserData("Enter your birthdate: ")
	var currUser *userhooks.UserDetails
	currUser, err :=userhooks.CreateNewUser(firstName, lastName, birthDate)
	if err!= nil{
		fmt.Println(err)
		return
	}
	admin := userhooks.CreateNewAdmin(randomdata.Email(), randomdata.Letters(6))

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputAdminCredentials()
	admin.OutputUserDetails()
	currUser.OutputUserDetails()
	currUser.ClearUserName()
	currUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	var userRes string
	fmt.Print(promptText)
	fmt.Scanln(&userRes)
	return userRes
}
