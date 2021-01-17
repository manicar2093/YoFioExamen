package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	LogTrace   *log.Logger
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

const logFileName = "logs.log"

func init() {

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileName, err))
	}

	LogTrace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogInfo = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogWarning = log.New(io.MultiWriter(logFile, os.Stdout),
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogError = log.New(io.MultiWriter(logFile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
