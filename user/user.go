package user

/* This is a whole new package for user, in this package we will have all the user related functions, variables, classes and structs.
We create packages like this to keep our code organized and easy to understand.*/
import (
	"errors"
	"fmt"
	"time"
)

// Creating a struct "globally"
// --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Here we will create a method for the struct "user"
/*
	Here, we will be passing "struct values" as arguments inside this function.

We will create a parameter and give it a name (u) and this parameter will be of type (user) which is the struct. This function "OutputUserData()" we are using to output the values of a specific user ("u")
*/
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
	// Outputting user details by applying a function into the code.
}

/*Here, we will create a mutation method. This method will be used to update the user details. In this method, inside the function, we will be using the "u" parameter which is of type "user"
But first we have to pass an "Asterisk" before the receiver parameter "u" to indicate that this method is a mutation method. */

// Creating the mutation method

func (u *User) UpdateUserDetails() {
	u.firstName = ""
	u.lastName = ""
}

// Here, we are going to create a constrcutor function to create a new Admin
func NewAdmin(email, password string) AdminUser {
	return AdminUser{
		email:    email,
		password: password,
		User: User{
			firstName: "Omar",
			lastName:  "Tamer",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}
}

// Here, we are going to create a constructor function to create a new user. This function will be used to create a new user. Usually by convention , the name of the constructor function is "New" followed by the name of the struct.
func New(FirstName, LastName, birthDate string) (*User, error) {
	// Here, we are goimg to create a validation struct. This struct will be used to validate the user details.
	if FirstName == "" || LastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name and Birth date are required!")
	}
	return &User{
		firstName: FirstName,
		lastName:  LastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time // Here, we called the "time" package, called it and created a nested struct out of it.
}

// Creating a new struct that inherits from the "User" struct and we will call it "AdminUser"
type AdminUser struct {
	email    string
	password string
	User     User
}
