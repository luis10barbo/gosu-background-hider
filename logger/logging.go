package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	fatalLogger	  *log.Logger
)

func init() {
	file, err := os.OpenFile("background_remover.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	fatalLogger = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InfoLog(message ...interface{}) {
	fmt.Println(message...)
	infoLogger.Println(message...)
}

func WarnLog(message ...interface{}) {
	fmt.Println(message...)
	warningLogger.Println(message...)
}

func ErrorLog(message ...interface{}) {
	fmt.Println(message...)
	errorLogger.Println(message...)
}

func FatalLog(message ...interface{}) {
	fmt.Println(message...)
	fatalLogger.Fatal(message...)
}