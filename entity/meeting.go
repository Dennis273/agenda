package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type Meeting struct {
	Title        string
	Holder       string
	Participator []string
	StartTime    int64
	EndTime      int64
}

const meetingFilePath string = "./meeting.json"

func CreateMeeting(meeting Meeting) (bool, string) {
	allMeetings := readMeetingsFromFile()
	allUsers := readUsersFromFile()
	currentUser := readCurrentUserFromFile()
	meeting.Holder = currentUser
	for _, temp := range allMeetings {
		if temp.Title == meeting.Title {
			return false, "meeting exit"
		}
	}
	for _, user1 := range meeting.Participator {
		var flag = true
		for _, user2 := range allUsers {
			if user1 == user2.Username {
				flag = false
				break
			}
		}
		if flag {
			return false, fmt.Sprintf("%s is not exit", user1)
		}
	}
	for _, temp := range allMeetings {
		if (meeting.StartTime < temp.StartTime && meeting.EndTime > temp.StartTime) || (meeting.StartTime < temp.EndTime && meeting.EndTime > temp.EndTime) {
			if temp.Holder == meeting.Holder {
				return false, "You are busy at that time"
			}
			for _, user1 := range meeting.Participator {
				if user1 == meeting.Holder {
					return false, fmt.Sprintf("%s is busy at that time", user1)
				}
				for _, user2 := range temp.Participator {
					if meeting.Holder == user2 {
						return false, "You are busy at that time"
					}
					if user1 == user2 {
						return false, fmt.Sprintf("%s is busy at that time", user1)
					}
				}
			}
		}
	}
	return true, ""
}

func AddMemberToMeeting(title string, user string) (bool, string) {
	return true, ""
}

func RemoveMemberFromMeeting(title string, user string) (bool, string) {
	return true, ""
}

func QueryMeeting(startTime, endTime int64) []Meeting {
	return make([]Meeting, 0)
}

func CancelMeeting(title string) (bool, string) {
	return true, ""
}

func QuitMeeting(title string) (bool, string) {
	return true, ""
}

func ClearMeeting() (bool, string) {
	return true, ""
}

func readMeetingsFromFile() []Meeting {
	meetings := make([]Meeting, 0)
	file, err := os.Open(meetingFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fInfo, err := os.Stat(meetingFilePath)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, fInfo.Size())
	_, err = file.Read(buffer)
	if fInfo.Size() == 0 {
		return meetings
	}
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(buffer, meetings)
	if err != nil {
		panic(err)
	}
	return meetings
}

func writeMeetingsIntoFile(meetings []Meeting) {
	data, err := json.Marshal(meetings)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(meetingFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
