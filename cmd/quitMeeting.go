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
	"os"

	"github.com/spf13/cobra"
)

var title string

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting -t [meeting title]",
	Short: "Quit a meeting that you have participated in",
	Long:  "Quit a meeting that you have participated in, when there is none participator in the meeting, the meeting will be cancelled ",
	Run: func(cmd *cobra.Command, args []string) {

		if title != "" {
			fmt.Printf("You will quit from %s meeting\n", title)
		} else {
			fmt.Fprintf(os.Stderr, "You have to input title by -t\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(quitMeetingCmd)
	quitMeetingCmd.PersistentFlags().StringVarP(&(title), "title", "t", "", "meeting title")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
