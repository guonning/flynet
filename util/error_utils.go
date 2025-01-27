package util

import (
	"fmt"
	log2 "github.com/asche910/flynet/log"
	"log"
)

var logger *log.Logger

func InitLog()  {
	logger = log2.GetLogger()
}

// just check error and print if err is not nil
func CheckError(err error, info string) {
	if err != nil {
		logger.Println(info, err)
	}
}

// check error and exit if err is not nil
func CheckErrorOrExit(err error, info string) {
	if err != nil {
		logger.Panicln(info, err)
	}
}

// get info about port occupied
func PortOccupiedInfo(port string) string {
	return fmt.Sprintf("port %s has been occuried!", port)
}

func AcceptErrorInfo() string {
	return "accept client error!"
}
