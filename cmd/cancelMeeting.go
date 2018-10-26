// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

var cancelMeetingArgs struct {
	title string
}
// cancelMeetingCmd represents the cancelMeeting command
var cancelMeetingCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "Cancel a meeting that you have created",
	Long:  "Cancel a meeting that you have created, when there are participators in the meeting, they will quit from the meeting",
	Run: func(cmd *cobra.Command, args []string) {
		if cancelMeetingArgs.title != "" {
			fmt.Printf("You will cancel %s meeting\n", cancelMeetingArgs.title)
		} else {
			fmt.Println("You have to input title by -t\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelMeetingCmd)
	cancelMeetingCmd.PersistentFlags().StringVarP(&(cancelMeetingArgs.title), "title", "t", "", "meeting title")
}