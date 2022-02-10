package logger

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

const (
	Reset = "\033[0m"

	PrimaryRed    = "\033[1;41m"
	PrimaryGreen  = "\033[1;42m"
	PrimaryYellow = "\033[1;43m"
	PrimaryCyan   = "\033[1;46m"

	SecondaryRed    = "\033[0;91m"
	SecondaryGreen  = "\033[0;92m"
	SecondaryYellow = "\033[0;93m"
	SecondaryCyan   = "\033[0;96m"
)

var (
	Limit                int
	LogLevel, LastUpdate string
	TimeZone             time.Location
	Style                bool
)

type Logger interface {
	Error(message string,args ...interface{})
	Success(message string,args ...interface{})
	Warning(message string,args ...interface{})
	Info(message string,args ...interface{})
}

// L is the global instance of the logger
var L = &StdOutLogger{}

// StdOutLogger logs to standard out
type StdOutLogger struct{}

func NewStdOutLogger(limit int, logLevel, timeZone string, style bool) Logger {
	var (
		newLine               string
		availabilityLogFolder bool = false
		availabilitylogLevel  bool = false
	)

	for _, value := range []string{"all", "error", "success", "warning", "info"} {
		if value == logLevel {
			availabilitylogLevel = true
		}
	}

	if !availabilitylogLevel {
		panic(errors.New("unknown log level " + logLevel))
	}

	timeLocation, err := time.LoadLocation(timeZone)
	if err != nil {
		panic(err)
	}

	Limit = limit
	LogLevel = logLevel
	TimeZone = *timeLocation
	Style = style

	if runtime.GOOS == "windows" {
		newLine = "\r\n"
	} else {
		newLine = "\n"
	}

	ignoreFile, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer ignoreFile.Close()

	contentIgnoreFile, err := ioutil.ReadFile(".gitignore")
	if err != nil {
		panic(err)
	}

	contents := strings.Split(string(contentIgnoreFile), newLine)

	for i := 0; i < len(contents); i++ {
		if contents[i] == "logs/" {
			availabilityLogFolder = true
		}
	}

	if !availabilityLogFolder {
		ignoreFile.Write([]byte("logs/" + newLine))
	}

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	return StdOutLogger{}
}

func (s StdOutLogger) ExecutionLimit() {
	CurrentDate := time.Now().In(&TimeZone).Format("2006-01-02")

	if LastUpdate != CurrentDate {
		if logFiles, _ := filepath.Glob("logs/*"); len(logFiles) > Limit+1 {
			sort.Strings(logFiles)

			os.Remove(logFiles[0])

			LastUpdate = CurrentDate
		}
	}
}

func (s StdOutLogger) Error(message string,args ...interface{}) {
	message = fmt.Sprintf(message+"\n", args...)

	if LogLevel == "all" || LogLevel == "error" {
		CurrentDate := time.Now().In(&TimeZone).Format("2006-01-02")

		logFile, _ := os.OpenFile("logs/"+CurrentDate+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		errorLog := log.New(logFile, "[ ERROR ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Llongfile)
		errorLog.Println(message)

		go s.ExecutionLimit()
	}

	var CurrentDatetime = time.Now().In(&TimeZone).Format("2006-01-02 15:04:05")

	if Style {
		fmt.Println(CurrentDatetime + " " + PrimaryRed + "[ ERROR ]" + SecondaryRed + " " + message + Reset)
	} else {
		fmt.Println(CurrentDatetime + " [ ERROR ] " + message)
	}
}

func (s StdOutLogger) Success(message string,args ...interface{}) {
	message = fmt.Sprintf(message+"\n", args...)
	if LogLevel == "all" || LogLevel == "success" {
		CurrentDate := time.Now().In(&TimeZone).Format("2006-01-02")

		logFile, _ := os.OpenFile("logs/"+CurrentDate+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		successLog := log.New(logFile, "[ SUCCESS ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Llongfile)
		successLog.Println(message)

		go s.ExecutionLimit()
	}

	var CurrentDatetime = time.Now().In(&TimeZone).Format("2006-01-02 15:04:05")

	if Style {
		fmt.Println(CurrentDatetime + " " + PrimaryGreen + "[ SUCCESS ]" + SecondaryGreen + " " + message + Reset)
	} else {
		fmt.Println(CurrentDatetime + " [ SUCCESS ] " + message)
	}
}

func (s StdOutLogger) Warning(message string,args ...interface{}) {
	message = fmt.Sprintf(message+"\n", args...)
	if LogLevel == "all" || LogLevel == "warning" {
		CurrentDate := time.Now().In(&TimeZone).Format("2006-01-02")

		logFile, _ := os.OpenFile("logs/"+CurrentDate+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		warningLog := log.New(logFile, "[ WARNING ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Llongfile)
		warningLog.Println(message)

		go s.ExecutionLimit()
	}

	var CurrentDatetime = time.Now().In(&TimeZone).Format("2006-01-02 15:04:05")

	if Style {
		fmt.Println(CurrentDatetime + " " + PrimaryYellow + "[ WARNING ]" + SecondaryYellow + " " + message + Reset)
	} else {
		fmt.Println(CurrentDatetime + " [ WARNING ] " + message)
	}
}

func (s StdOutLogger) Info(message string,args ...interface{}) {
	message = fmt.Sprintf(message+"\n", args...)
	if LogLevel == "all" || LogLevel == "info" {
		CurrentDate := time.Now().In(&TimeZone).Format("2006-01-02")

		logFile, _ := os.OpenFile("logs/"+CurrentDate+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		infoLog := log.New(logFile, "[ INFO ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Llongfile)
		infoLog.Println(message)

		go s.ExecutionLimit()
	}

	var CurrentDatetime = time.Now().In(&TimeZone).Format("2006-01-02 15:04:05")

	if Style {
		fmt.Println(CurrentDatetime + " " + PrimaryCyan + "[ INFO ]" + SecondaryCyan + " " + message + Reset)
	} else {
		fmt.Println(CurrentDatetime + " [ INFO ] " + message)
	}
}