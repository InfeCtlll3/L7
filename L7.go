package L7

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// define all the constants for the code
const (
	NoTime        = 1
	FullTimeStamp = 2
	Zulu          = 3
	Epoch         = 4
	CRITICAL      = 50
	ERROR         = 40
	WARNING       = 30
	INFO          = 20
	DEBUG         = 10
)

// LoggerStruct defines the structure of the logging object.
type LoggerStruct struct {
	logTimeStamp int
	logLevel     int
}

// Params is the structure used to create the Logger object
type Params struct {
	TimeStampFormat, LogLevel int
}

// validateLogLevel will validate if the logging level is valid.
func validateLogLevel(loggingLevel int) bool {
	if loggingLevel%10 == 0 && loggingLevel >= DEBUG && loggingLevel <= CRITICAL {
		return true
	}
	return false
}

// Logger is a method to create the logging object to be used by Log
// will return a struct of LoggerStruct
// This method is used in order to set defaults
func Logger(kargs Params) LoggerStruct {
	if kargs.TimeStampFormat == 0 {
		kargs.TimeStampFormat = FullTimeStamp
	}
	if kargs.LogLevel == 0 {
		kargs.LogLevel = ERROR
	}
	if !validateLogLevel(kargs.LogLevel) || kargs.TimeStampFormat > Epoch || kargs.TimeStampFormat < NoTime {
		panic("Misconfiguration when configuring the Logger object.")
	} else {
		return LoggerStruct{kargs.TimeStampFormat,
			kargs.LogLevel}
	}
}

// Log is the method used to log you messages.
func (context *LoggerStruct) Log(messageLevel int, message string) {
	if messageLevel >= context.logLevel {
		var currentTime, currentLevel string

		// define the log context
		switch context.logTimeStamp {
		case NoTime:
			// case 1 means no timestamp
			currentTime = ""
		case FullTimeStamp:
			// case 2 means full local timestamp
			currentTime = time.Now().Format("15:04:05 02/01/2006")
		case Zulu:
			// case 3 means zulu format
			currentTime = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		case Epoch:
			// case 4 means epoch
			currentTime = strconv.Itoa(int(time.Now().Unix()))
		}
		// define a nicely verbose loglevel name
		switch context.logLevel {
		case DEBUG:
			// case 1 means debug
			currentLevel = "[DEBUG]"
		case INFO:
			// case 2 means info
			currentLevel = "[INFO]"
		case WARNING:
			currentLevel = "[WARNING]"
		case ERROR:
			// case 4 means error
			currentLevel = "[ERROR]"
		case CRITICAL:
			// case 5 means critical
			currentLevel = "[CRITICAL]"
		}
		fmt.Println(currentTime, currentLevel, trace(), message)
	}
}

// SetLogLevel is the method used to change to different log levels
// May be useful to enable/disable logging at certain parts of the code
func (context *LoggerStruct) SetLogLevel(newLevel int) {
	if validateLogLevel(newLevel) {
		if context.logLevel != newLevel {
			context.logLevel = newLevel
		}
	} else {
		fmt.Println("Log level is not supported. Please specify a valid log level.")
	}
}

// trace is used internally to do profiling of the caller method
func trace() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "?"
	}

	fn := runtime.FuncForPC(pc)
	return "(" + fn.Name() + ")"
}
