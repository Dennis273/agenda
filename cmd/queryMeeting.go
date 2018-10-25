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

	"github.com/spf13/cobra"
)

var queryMeetingArgs struct {
	startTime string
	endTime   string
}

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
	Use:   "queryMeeting",
	Short: "query meetings with a time interval",
	Long:  `query meetings with a time interval`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("queryMeeting called")
		fmt.Printf("the start time: %s\n", queryMeetingArgs.startTime)
		fmt.Printf("the end time: %s", queryMeetingArgs.endTime)
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingCmd)

	queryMeetingCmd.Flags().StringVarP(&(queryMeetingArgs.startTime), "start", "s", "", "the start time")
	queryMeetingCmd.Flags().StringVarP(&(queryMeetingArgs.endTime), "end", "e", "", "the start time")
}
