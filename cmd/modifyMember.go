// Copyright © 2018 dengzijie
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

	"github.com/Dennis273/agenda/entity"
	"github.com/spf13/cobra"
)

var modifyMemberArgs struct {
	title     string
	addMember []string
	delMember []string
}

// modifyMemberCmd represents the modifyMember command
var modifyMemberCmd = &cobra.Command{
	Use:   "modifyMember",
	Short: "Add or delete member(s)",
	Long:  `Add member(s) to into speacified meeting or delete member(s) from it`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modifyMember called")
		fmt.Println("meeting title: ", modifyMemberArgs.title)
		fmt.Printf("new participators: ")
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("ModifyMeeting failed: %s", r)
				entity.Error(r)
			} else {
				fmt.Println("ModifyMeeting success")
			}
		}()
		for num, name := range modifyMemberArgs.addMember {
			fmt.Printf("%d:%s ", num, name)
		}
		fmt.Printf("\nold participators:")
		for num, name := range modifyMemberArgs.delMember {
			fmt.Printf("%d:%s ", num, name)
		}
		fmt.Print("\n")
		entity.Info("Modify Member called")
		entity.ModifyMeeting(modifyMemberArgs.title, modifyMemberArgs.addMember, modifyMemberArgs.delMember)
	},
}

func init() {
	rootCmd.AddCommand(modifyMemberCmd)

	modifyMemberCmd.Flags().StringVarP(&(modifyMemberArgs.title), "title", "t", "", "meeting title")
	modifyMemberCmd.Flags().StringArrayVarP(&(modifyMemberArgs.addMember), "add", "a", nil, "add member(s) to the meeting")
	modifyMemberCmd.Flags().StringArrayVarP(&(modifyMemberArgs.delMember), "delete", "d", nil, "delete member(s) from the meeting")
}
