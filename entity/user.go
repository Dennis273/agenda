package entity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}

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

// check username and write new user info into file
func CreateUser(user User) (bool, string) {
	return true, ""
}

// check password and write userinfo into currentUsr file
func Login(username, password string) (bool, string) {
	return true, ""
}

// clear userinfo in currentUsr file
func Logout() (bool, string) {
	return true, ""
}

// remove user info from file
func DeleteUser() (bool, string) {
	return true, ""
}

// query all userinfo from file
func QueryUser() []User {
	return make([]User, 0)
}

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
	err = json.Unmarshal(buffer, &users)
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
	file, err := os.OpenFile(userFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}

func readCurrentUserFromFile() (username string) {
	file, err := os.Open(currentUserFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file_reader := bufio.NewReader(file)
	username, err = file_reader.ReadString('\n')
	username = strings.Replace(username, "\n", "", -1)
	if err != nil {
		panic(err)
	}
	return username
}

func writeCurrentUserToFile(username string) {
	var str string
	file, err := os.OpenFile(currentUserFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	str = fmt.Sprintf("%s\n", username)
	_, err = file.WriteString(str)
	if err != nil {
		panic(err)
	}
}
