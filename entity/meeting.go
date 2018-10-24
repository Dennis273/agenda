package entity

import (
	"encoding/json"
	"os"
	"time"
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

}

func AddMemberToMeeting(title string, user string) (bool, string) {

}

func RemoveMemberFromMeeting(title string, user string) (bool, string) {

}

func QueryMeeting() []Meeting {

}

func CancelMeeting(title string) (bool, string) {

}

func QuitMeeting(title string) (bool, string) {

}

func ClearMeeting() (bool, string) {

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
