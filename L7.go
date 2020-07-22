package L7

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
	"os"
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
	Console = false
	Logfile = true
)

// LoggerStruct defines the structure of the logging object.
type LoggerStruct struct {
	logTimeStamp int
	logLevel     int
	logStdOut	 bool
	logFileName string
}

// Params is the structure used to create the Logger object
type Params struct {
	TimeStampFormat, LogLevel  int
	LogFileName string
	LogToConsoleOrFile bool
}

func GetEnv(env string) int {
	if os.Getenv(env) != "" {
		toreturn, _ := strconv.Atoi(env)
		return toreturn
	}
}

// validateLogLevel will validate if the logging level is valid.
func validateLogLevel(loggingLevel int) bool {
	if loggingLevel%10 == 0 && loggingLevel >= DEBUG && loggingLevel <= CRITICAL {
		return true
	}
	return false
}

func stringBuilder(Strings []string) string {
	if len(Strings) == 0 {
		return ""
	} else {
		if len(Strings) == 1 {
			return Strings[0]
		}
	}
	var strReturn strings.Builder
	for _, str := range Strings {
		strReturn.WriteString(str)
	}
	return strReturn.String()
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
	if kargs.LogStdOutDst == false {
		kargs.LogStdOutDst = Console
	}
	if !validateLogLevel(kargs.LogLevel) || kargs.TimeStampFormat > Epoch || kargs.TimeStampFormat < NoTime ||
		kargs.LogFileName == "" {
		panic("Misconfiguration when configuring the Logger object.")
	} else {
		return LoggerStruct{kargs.TimeStampFormat,
			kargs.LogLevel, kargs.LogStdOutDst, kargs.LogFileName}
	}
}

func (context *LoggerStruct) LogErrorIfAny(err error, messages ...string) {
	messageLevel := ERROR
	if err != nil {
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
			switch messageLevel {
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
			if context.logStdOut == Console {
				fmt.Println(currentTime, currentLevel, trace(), stringBuilder(append(messages, err.Error())))
			} else {
				dumpLogToFile(context.logFileName, currentTime, currentLevel, trace(),
					stringBuilder(append(messages, err.Error())))
			}
		}
	}
}

// Log is the method used to log you messages.
func (context *LoggerStruct) Log(messageLevel int, messages ...string) {
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
		switch messageLevel {
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
		if context.logStdOut == Console {
			fmt.Println(currentTime, currentLevel, trace(), stringBuilder(messages))
		} else {
			dumpLogToFile(context.logFileName, currentTime, currentLevel, trace(), stringBuilder(messages))
		}
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
	return stringBuilder([]string{"(", fn.Name(), ")"})
}

func dumpLogToFile(kargs...string) {
	f, err := os.OpenFile(kargs[0],
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	space := " "
	if _, err := f.WriteString(stringBuilder([]string{kargs[1], space, kargs[2], space, kargs[3], space,
		kargs[4], "\n"})); err != nil {
		fmt.Println(err)
	}
}