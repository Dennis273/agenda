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
	"regexp"

	"math"

	"github.com/Dmaxzj/agenda/entity"
	"github.com/spf13/cobra"
)

var createMeetingArgs struct {
	meetingTitle  string
	startTime     string
	endTime       string
	participators []string
}

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "Create a meeting",
	Long:  "Create a meeting as sponsor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createMeeting called")
		fmt.Println("meeting title: ", createMeetingArgs.meetingTitle)
		fmt.Println("start time: ", createMeetingArgs.startTime)
		fmt.Println("end time: ", createMeetingArgs.endTime)
		fmt.Printf("participators: \n")
		for num, name := range createMeetingArgs.participators {
			fmt.Printf("%d:%s \n", num, name)
		}

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Create meeting failed:", r)
			} else {
				fmt.Println("Create meeting success")
			}
		}()

		if !checkTimeFormat(createMeetingArgs.startTime) || !checkTimeFormat(createMeetingArgs.endTime) {
			panic(fmt.Sprint("Invalid date"))
		}
		var year, month, day, hour, min int
		var stime, etime int64
		fmt.Sscanf(createMeetingArgs.startTime, "%d-%d-%d-%d-%d", &year, &month, &day, &hour, &min)
		stime = int64(year*int(math.Pow10(8)) + month*int(math.Pow10(6)) + day*int(math.Pow10(4)) + hour*int(math.Pow10(2)) + min)
		fmt.Sscanf(createMeetingArgs.endTime, "%d-%d-%d-%d-%d", &year, &month, &day, &hour, &min)
		etime = int64(year*int(math.Pow10(8)) + month*int(math.Pow10(6)) + day*int(math.Pow10(4)) + hour*int(math.Pow10(2)) + min)
		if stime >= etime {
			panic("time interval error")
		}
		new_meeting := entity.Meeting{createMeetingArgs.meetingTitle,
			"",
			createMeetingArgs.participators,
			stime,
			etime}
		success, err := entity.CreateMeeting(new_meeting)
		if !success {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringVarP(&(createMeetingArgs.meetingTitle), "title", "t", "", "meeting title")
	createMeetingCmd.Flags().StringVarP(&(createMeetingArgs.startTime), "startTime", "s", "", "the time when the meeting begin(yyyy-MM-dd-hh-mm)")
	createMeetingCmd.Flags().StringVarP(&(createMeetingArgs.endTime), "endTime", "e", "", "the end time of the meeting(yyyy-MM-dd-hh-mm)")
	createMeetingCmd.Flags().StringArrayVarP(&(createMeetingArgs.participators), "participators", "p", nil, "the participators of the meeting")
}

func checkTimeFormat(time string) bool {
	var year, month, day, hour, min int
	if matched, _ := regexp.MatchString("[\\d]{4}-[\\d]{2}-[\\d]{2}-[\\d]{2}-[\\d]{2}", time); !matched {
		panic("date formate error")
	}
	fmt.Sscanf(createMeetingArgs.startTime, "%d-%d-%d-%d-%d", &year, &month, &day, &hour, &min)
	if min < 0 || min >= 60 || hour < 0 || hour >= 24 || day >= 32 || day <= 0 || month <= 0 || month > 12 || year < 0 {
		return false
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return true
	case 4, 6, 9, 11:
		if day == 31 {
			return true
		}
	default:
		if day <= 28 {
			return true
		} else {
			return year%400 == 0 || (year%4 == 0 && year%100 != 0)
		}
	}
	return true
}
