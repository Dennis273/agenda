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
	"github.com/dashuibihello/agenda/entity"
	"github.com/spf13/cobra"
)

var quitMeetingArgs struct {
	title string
}

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting -t [meeting title]",
	Short: "Quit a meeting that you have participated in",
	Long:  "Quit a meeting that you have participated in. When there is none participator in the meeting, the meeting will be cancelled ",
	Run: func(cmd *cobra.Command, args []string) {
		entity.QuitMeeting(quitMeetingArgs.title)
	},
}

func init() {
	rootCmd.AddCommand(quitMeetingCmd)
	quitMeetingCmd.PersistentFlags().StringVarP(&(quitMeetingArgs.title), "title", "t", "", "meeting title")
}
