package userhooks


import (
	"errors"
	"fmt"
	"time"
)


type UserDetails struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type AdminDetails struct{
	email string
	password string
	UserDetails

}
func (u *UserDetails) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
func (u *UserDetails) ClearUserName() {
	u.firstName = ""
	u.lastName = ""

}
func CreateNewUser (firstName, lastname, birthDate string) (*UserDetails, error){
	if firstName == "" || lastname == "" || birthDate == ""{
		return nil, errors.New("firstname, lastname and birthdate are required")
	}
	return &UserDetails{
		firstName: firstName,
		lastName: lastname,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (a AdminDetails) OutputAdminCredentials(){
	fmt.Println(a.email, a.password)

}

func CreateNewAdmin(email , password string) AdminDetails{
	return AdminDetails{
		email: email,
		password: password,
		UserDetails: UserDetails{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}

}