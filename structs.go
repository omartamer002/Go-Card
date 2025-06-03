package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	UserFirstName := getUserData("Please enter your first name: ")
	UserLastName := getUserData("Please enter your last name: ")
	Userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(UserFirstName, UserLastName, Userbirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	/* Here we created a variable called "appUser" of type "user" struct. We created this variable that takes the instance value of the struct.
	This means by creating this variable "appUser" we will be equalizing the vaule (components) of the struct and the variables (firstName, lastName, birthdate)*/

	admin := user.NewAdmin("omarscareer002@gmail.com", "12345678")
	admin.User.OutputUserDetails()
	admin.User.UpdateUserDetails()
	admin.User.OutputUserDetails()

	appUser.OutputUserDetails() // Here we are calling the method "OutputUserDetails()" and passing the "appUser". This is for outputting user details by applying a function into the code.
	appUser.UpdateUserDetails() // Here we are calling the method "UpdateUserDetails()" and passing the "appUser". This is for updating user details by applying a function into the code.
	appUser.OutputUserDetails() /* Hesre we are calling the method "OutputUserDetails()" and passing the "appUser". This is for outputting user details after updating and applying mutation method by applying a function into the code.*/

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
