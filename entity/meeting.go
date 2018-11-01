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

var meetingData []Meeting

const meetingFilePath string = "./meeting.json"

/*
func CreateMeeting(meeting Meeting) (bool, string) {

}

func AddMemberToMeeting(title string, user string) (bool, string) {

}

func QueryMeeting() []Meeting {

}
*/

func RemoveMemberFromMeeting(title string, user string) bool {
	if CheckUserLogin() != true {
		fmt.Println("You have to log in first!")
		return false
	}
	meeting, flag, i := QueryMeetingbyTitle(title)
	if flag == false {
		fmt.Println("The meeting dose not exist!")
		return false
	}
	for j := 0; j < len(meetingData[j].Participator); j++ {
		if meetingData[i].Participator[j] == User_Login.Username {
			meetingData[i].Participator = append(meetingData[i].Participator[:j], meetingData[i].Participator[j+1:]...)
			fmt.Println("Remove participator successfully!")
			if len(meetingData[j].Participator) == 0 {
				fmt.Println("Becasue there is no participator in this meeting, trying to cancel the meeting!")
				CancelMeeting(meeting.Title)
			}
			return true
		}
	}
	fmt.Println("The meeting don't have this participator!")
	return false
}

func QueryMeetingbyTitle(title string) (Meeting, bool, int) {
	for i := 0; i < len(meetingData); i++ {
		if meetingData[i].Title == title {
			return meetingData[i], true, i
		}
	}
	var tmp Meeting
	return tmp, false, 0
}

func CancelMeeting(title string) bool {
	if CheckUserLogin() != true {
		fmt.Println("You have to log in first!")
		return false
	}
	meeting, flag, i := QueryMeetingbyTitle(title)
	if flag == false {
		fmt.Println("The meeting dose not exist!")
		return false
	}
	if meeting.Holder != User_Login.Username {
		return false
	}
	meetingData = append(meetingData[:i], meetingData[i+1:]...)
	fmt.Println("The meeting is cancelled sucessfuly!")
	return true

}

func QuitMeeting(title string) bool {
	if CheckUserLogin() != true {
		fmt.Println("You have to log in first!")
		return false
	}
	if RemoveMemberFromMeeting(title, User_Login.Username) == true {
		fmt.Println("Your quit the meeting successfully")
		return true
	}
	return false
}

func ClearMeeting() bool {
	count := -1
	if CheckUserLogin() != true {
		fmt.Println("You have to log in first!")
		return false
	}
	for i := 0; i < len(meetingData); i++ {
		if meetingData[i].Holder == User_Login.Username {
			if CancelMeeting(meetingData[i].Title) {
				count++
			}
		}
	}
	if count == 0 {
		fmt.Println("You don't have any meeting!")
		return false
	}
	fmt.Println("You clear all your meetings sucessfully!")
	return true
}

func readMeetingsFromFile() []Meeting {
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
	if err != nil {
		panic(err)
	}
	meetings := make([]Meeting, 0)
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
