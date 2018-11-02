package entity

import (
	"log"
	"os"
)

var logName = "./agenda.log"

func Info(msg string) {
	logFile, err1 := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	defer logFile.Close()
	if err1 != nil {
		log.Fatalln("open log file error !")
	}
	Log := log.New(logFile, "[INFO]", log.LstdFlags)
	Log.Println(msg)
}

func Error(msg interface{}) {
	logFile, err1 := os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	defer logFile.Close()
	if err1 != nil {
		log.Fatalln("open log file error !")
	}
	Log := log.New(logFile, "[ERROR]", log.LstdFlags)
	Log.Println(msg)
}
