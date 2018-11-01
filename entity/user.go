package entity

import (
	"encoding/json"
	"os"
)

type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}

var login_flag bool

var User_Login User

const userFilePath string = "./user.json"
const currentUserFilePath string = "./currentUser.txt"

func init() {
	// check if user.json exist
	if _, err := os.Stat(userFilePath); os.IsNotExist(err) {
		_, err := os.Create(userFilePath)
		if err != nil {
			panic(err)
		}
	}
	//check if currentUser.txt exist
	if _, err := os.Stat(currentUserFilePath); os.IsNotExist(err) {
		_, err := os.Create(currentUserFilePath)
		if err != nil {
			panic(err)
		}
	}
}

// check if the user has logged in
func CheckUserLogin() bool {
	login_flag = true
	return login_flag
}

/*
// check username and write new user info into file
func CreateUser(user User) (bool, string) {

}

// check password and write userinfo into currentUsr file
func Login(username, password string) (bool, string) {

}

// clear userinfo in currentUsr file
func Logout() (bool, string) {

}

// remove user info from file
func DeleteUser() (bool, string) {

}

// query all userinfo from file
func QueryUser() []User {

}

*/

func readUsersFromFile() []User {
	file, err := os.Open(userFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fInfo, err := os.Stat(userFilePath)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, fInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
	users := make([]User, 0)
	err = json.Unmarshal(buffer, users)
	if err != nil {
		panic(err)
	}
	return users
}

func writeUsersIntoFile(users []User) {
	data, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(userFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
