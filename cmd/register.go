// Copyright © 2018 Dennis273 <dennic695@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var registerArgs struct {
	username string
	password string
	email string
	phone string
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register new user",
	Long:  `Register new user`,
	Run: func(cmd *cobra.Command, args []string) {
		p0 := ValidateUsername(registerArgs.username)
		p1 := ValidatePassword(registerArgs.password)
		p2 := ValidateEmail(registerArgs.email)
		p3 := ValidatePhoneNumber(registerArgs.phone)
		if !p0 || !p1 || !p2 || !p3 {
			fmt.Println("Invalid arguments => aborted.")
			return
		}
		fmt.Printf("%s %s %s %s\n", registerArgs.username, registerArgs.password,registerArgs.email, registerArgs.phone)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.PersistentFlags().StringVarP(&(registerArgs.username), "username", "u", "", "Username to login")
	registerCmd.PersistentFlags().StringVarP(&(registerArgs.password), "password", "p", "", "Password")
	registerCmd.PersistentFlags().StringVarP(&(registerArgs.email), "email", "e", "", "Email")
	registerCmd.PersistentFlags().StringVarP(&(registerArgs.phone), "phoneNumber", "n", "", "Phone number")
}