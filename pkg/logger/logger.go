package logger

import (
	"clean-architecture-beego/pkg/helpers/converter_value"
	"encoding/json"
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

	ERROR = "<ERROR>"
	INFO = "<INFO>"
	DEBUG = "<DEBUG>"

	XmodeTest = "Test"
	XmodeRunning = "Running"
)

var (
	Limit                int
	LogLevel, LastUpdate string
	TimeZone             time.Location
	Style                bool
	Version string
	App string
	Service string
	Xmode string
)


type Logger interface {
	Error(logging LoggingObj)
	Info(logging LoggingObj)
	Debug(logging LoggingObj)
}

// L is the global instance of the logger
var L = &StdOutLogger{}

// StdOutLogger logs to standard out
type StdOutLogger struct{}

func NewStdOutLogger(limit int, logLevel, timeZone string, style bool,version string,app string,service string,xmode string) Logger {
	Version = version
	App = app
	Service = service
	xmode = xmode
	var (
		newLine               string
		availabilityLogFolder bool = false
		availabilitylogLevel  bool = false
	)

	for _, value := range []string{"all", "error", "success", "warning", "info", "debug"} {
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

	if xmode != XmodeTest {
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

func (s StdOutLogger) Error(logging LoggingObj) {
	logData := s.generatePrefixLog(ERROR,logging,false)
	fmt.Println(logData)
	if (LogLevel == "all" || LogLevel == "error") && Xmode != XmodeTest{
		logData = s.generatePrefixLog(ERROR,logging,true)
		s.generateFile(logging.Feature,logData)
	}

}

func (s StdOutLogger) Info(logging LoggingObj) {
	logData := s.generatePrefixLog(INFO,logging,false)
	fmt.Println(logData)
	if (LogLevel == "all" || LogLevel == "info") && Xmode != XmodeTest {
		logData = s.generatePrefixLog(INFO,logging,true)
		s.generateFile(logging.Feature,logData)
	}

}

func (s StdOutLogger) Debug(logging LoggingObj) {
	logData := s.generatePrefixLog(DEBUG,logging,false)
	fmt.Println(logData)
	if (LogLevel == "all" || LogLevel == "debug") && Xmode != XmodeTest {
		logData = s.generatePrefixLog(DEBUG,logging,true)
		s.generateFile(logging.Feature,logData)
	}

}

func (s StdOutLogger) generatePrefixLog(loglevel string ,logging LoggingObj,toFile bool) string {
	var logString string
	jsonData ,_ := json.Marshal(logging.Data)
	now := converter_value.DateTimeToStringWithFormat(time.Now(),converter_value.DateTimeFormatDefault)
	if Style && !toFile{
		switch loglevel {
		case ERROR:
			logString =	fmt.Sprintln( PrimaryRed + loglevel + SecondaryRed + Version + " " + now + " " + logging.Host + " " + App + " " +
				logging.PathFile + " " + logging.RequestId + " " + string(jsonData) + " " + logging.Message + Reset)
		case INFO:
			logString =	fmt.Sprintln( PrimaryCyan + loglevel + SecondaryCyan + Version + " " + now + " " + logging.Host + " " + App + " " +
				logging.PathFile + " " + logging.RequestId + " " + string(jsonData) + " " + logging.Message + Reset)
		case DEBUG:
			logString =	fmt.Sprintln( PrimaryYellow + loglevel + SecondaryYellow + Version + " " + now + " " + logging.Host + " " + App + " " +
				logging.PathFile + " " + logging.RequestId + " " + string(jsonData) + " " + logging.Message + Reset)
		}
	} else {
		logString =	fmt.Sprintln(loglevel + Version + " " + now + " " + logging.Host + " " + App + " " +
			logging.PathFile + " " + logging.RequestId + " " + string(jsonData) + " " + logging.Message)
	}


	return logString
}

func (s StdOutLogger) generateFile(feature string,message string) {
	CurrentDate := time.Now().In(&TimeZone).Format("2006-01-02")
	name := App + "-" + Service + "." + feature + "." + CurrentDate
	logFile, _ := os.OpenFile("logs/"+name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()


	infoLog := log.New(logFile, "", log.Ldate|log.Ltime|log.Lmsgprefix|log.Llongfile)
	infoLog.SetFlags(0)
	infoLog.Println(message)

	go s.ExecutionLimit()
}

