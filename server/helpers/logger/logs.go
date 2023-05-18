package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	LogsDirpath = "../logs"
)

func Init() {

	err := os.Mkdir(LogsDirpath, 0777)
	if err != nil {
		if os.IsExist(err) {
			log.Println("file already exists")
		}
	}
}

func SetLogFile() *os.File {
	year, month, day := time.Now().Date()
	fileName := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
	filePath, err := os.OpenFile(LogsDirpath+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Print(err)
	}
	return filePath
}

func Error(args ...interface{}) {
	var msg []interface{}
	getFilePath := SetLogFile()
	pre := log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	for _, value := range args {
		switch v := value.(type) {
		default:
			msg = append(msg, v)
		}
	}
	pre.Print(msg)
}
func Info(args ...interface{}) {
	var msg []interface{}
	// getFilePath := SetLogFile()
	// pre := log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	for _, value := range args {
		switch v := value.(type) {
		default:
			msg = append(msg, v)
		}
	}
	log.Print("INFO: ", log.Ldate|log.Ltime|log.Lshortfile, msg)
}

func Warning(args ...interface{}) {
	var msg []interface{}
	getFilePath := SetLogFile()
	pre := log.New(getFilePath, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	for _, value := range args {
		switch v := value.(type) {
		default:
			msg = append(msg, v)
		}
	}
	pre.Print(msg)
}
