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
	"math"
	"strings"

	"github.com/Dmaxzj/agenda/entity"
	"github.com/spf13/cobra"
)

var queryMeetingArgs struct {
	startTime string
	endTime   string
}

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
	Use:   "queryMeeting",
	Short: "Query meetings with a time interval which involves you.",
	Long:  `Query meetings with a time interval which involves you.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("queryMeeting called")
		fmt.Printf("the start time: %s\n", queryMeetingArgs.startTime)
		fmt.Printf("the end time: %s", queryMeetingArgs.endTime)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Query meeting failed:", r)
				entity.Error(r)
			} else {
				fmt.Println("Query meeting success")
			}
		}()
		if !checkTimeFormat(queryMeetingArgs.startTime) || !checkTimeFormat(queryMeetingArgs.endTime) {
			panic(fmt.Sprint("Invaild date"))
		}
		var year, month, day, hour, min int
		var stime, etime int64
		fmt.Sscanf(queryMeetingArgs.startTime, "%d-%d-%d-%d-%d", &year, &month, &day, &hour, &min)
		stime = int64(year*int(math.Pow10(8)) + month*int(math.Pow10(6)) + day*int(math.Pow10(4)) + hour*int(math.Pow10(2)) + min)
		fmt.Sscanf(queryMeetingArgs.endTime, "%d-%d-%d-%d-%d", &year, &month, &day, &hour, &min)
		etime = int64(year*int(math.Pow10(8)) + month*int(math.Pow10(6)) + day*int(math.Pow10(4)) + hour*int(math.Pow10(2)) + min)
		if stime > etime {
			panic("time interval error")
		}
		entity.Info("Query Meeting called")
		meetings := entity.QueryMeeting(stime, etime)
		var output strings.Builder
		output.WriteString("\nStartTime\tEndTime\t\tTitle\t\tHolder\tParticipators\n")
		for _, meeting := range meetings {
			str := fmt.Sprintf("%s\t%s\t%s\t%s\t", dateInt64ToString(meeting.StartTime), dateInt64ToString(meeting.EndTime), meeting.Title, meeting.Holder)
			output.WriteString(str)
			for _, p := range meeting.Participator {
				output.WriteString(p)
				output.WriteString("  ")
			}
			output.WriteString("\n")
		}
		fmt.Print(output.String())
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingCmd)

	queryMeetingCmd.Flags().StringVarP(&(queryMeetingArgs.startTime), "startTime", "s", "", "the start time(yyyy-MM-dd-hh-mm)")
	queryMeetingCmd.Flags().StringVarP(&(queryMeetingArgs.endTime), "endTime", "e", "", "the end time(yyyy-MM-dd-hh-mm)")
}

func dateInt64ToString(dateI int64) (date string) {
	var year, month, day, hour, min int64
	min = dateI % 100
	dateI /= 100
	hour = dateI % 100
	dateI /= 100
	day = dateI % 100
	dateI /= 100
	month = dateI % 100
	dateI /= 100
	year = dateI
	date = fmt.Sprintf("%v-%v-%v-%v-%v", year, month, day, hour, min)
	return date
}
