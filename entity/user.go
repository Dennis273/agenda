package entity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"io"
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
		file, err := os.OpenFile(userFilePath, os.O_WRONLY|os.O_TRUNC, 0777)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.Write([]byte("[]"))
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
func CreateUser(username, password, email, phone string) (bool, string) {
	users := readUsersFromFile()
	for _, value := range users {
		if value.Username == username {
			return false, "Username is owned."
		}
		if value.Email == email {
			return false, "Email is owned."
		}
	}
	user := User{
		username,
		password,
		email,
		phone,
	}
	writeUsersIntoFile(append(users, user))
	return true, ""
}

// check password and write userinfo into currentUsr file
func Login(username, password string) (bool, string) {
	cur := readCurrentUserFromFile()
	if cur != "" {
		return false, "Already logged in as " + cur
	}
	users := readUsersFromFile()
	for _, value := range users {
		if value.Username == username {
			if  value.Password == password {
				writeCurrentUserToFile(username)
				return true, ""
			} else {
				return false, "Password in correct."
			}
		}
	}
	return false, "User not found."
}

// clear userinfo in currentUsr file
func Logout() (bool, string) {
	if cur := readCurrentUserFromFile(); cur == "" {
		return false, "User not logged in."
	} else {
		writeCurrentUserToFile("")
		return true, ""
	}
}

// remove user info from file
func DeleteUser() (bool, string) {
	username := readCurrentUserFromFile()
	var i int
	if username == "" {
		fmt.Println("User not logged in")
		return false, "User not logged in"
	} else {
		users := readUsersFromFile()
		for index, value := range users {
			if value.Username == username {
				i = index
				break
			}
		}
		writeUsersIntoFile(append(users[:i], users[i + 1:]...))
		fmt.Println("Delete User success")
		return true, ""
	}
}

// query all userinfo from file
func QueryUser() (bool, string) {
	cur := readCurrentUserFromFile()
	if cur == "" {
		return false, "Permission denied"
	} else {
		users := readUsersFromFile()
		fmt.Println("Users:")
		for _, value := range users {
			fmt.Printf("%s\t%s\t%s\n", value.Username, value.Email, value.Phone)
		}
		return true, ""
	}
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
	if err != nil {
		if err ==  io.EOF {
			username = ""
		} else {
			panic(err)
		}
	} else {
		username = strings.Replace(username, "\n", "", -1)
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
