# L7
This project was created to do simple logging for my own projects. L7 was created to provide uniform logging providing multiple types of time stamps (local, zulu and epoch) and also tracing (e.g. which method logged this line).

### Features

- Pythonic way of logging data in golang.
- Support zulu, local and epoch time stamps.
- Uniform way of outputing data.
- Function context logging.
- Log levels.

### To-do
- Add support to write logs to a local file (using non-blocking goroutines).
- Add support to send logs to a proper log/syslog server.
- Add support to generate uuid(s) in order to make tracing easier.

### Contact

In case of problems, do not hesitate to cut a ticket or please e-mail me at contato.carmando@gmail.com

#### Installing

`$ go get github.com/InfeCtlll3/L7`



#### Example:　

```go
import (
	"github.com/InfeCtlll3/L7"
)

func main() {
	log := L7.Logger(L7.Params{TimeStampFormat: L7.Zulu, LogLevel: L7.DEBUG})
	// log level is DEBUG
	log.Log(L7.DEBUG, "This is a debug message!")
	log.SetLogLevel(L7.ERROR)
	// log level is now error
	log.Log(L7.DEBUG, "This message will not be displayed because of the log level")
	log.Log(L7.ERROR, "This is an error message!")
	// log also support multiple logging messages in a single call just like Println
	// This uses some performatic string builder (string buffers) in order to concatenate
	// This way you don't need to concatenate stuff when you call logging
	// Just make sure all the arguments are string :)
	veryImportantVariable := "Not so important"
	log.Log(L7.ERROR, "here is your variable: ", veryImportantVariable)
}
```

#### Output

```bash
2019-12-19T02:04:58.343Z [DEBUG] (main.main) This is a debug message!
2019-12-19T02:04:58.343Z [ERROR] (main.main) This is an error message!
2019-12-19T19:15:49.480Z [ERROR] (main.main) here is your variable: Not so important
```
#### Log Format
```
Timestamp [LogLevel] (context.method) error message	
```

### Supported logging levels
- Critical - `L7.CRITICAL`
- Error - `L7.ERROR`
- Warning - `L7.WARNING`
- Info - `L7.INFO`
- Debug - `L7.DEBUG`

### Supported time stamps
- No time stamps - `L7.NoTime`
- Host time - `L7.FullTimeStamp`
- Zulu - `L7.Zulu`
- Epoch - `L7.Epoch`

### Lib defaults
By default, whenever you create a Logger object with an empty Params struct (`log := L7.Logger(L7.Params{})`), it will use `L7.FullTimeStamp` for time stamps and `L7.ERROR`for the log level.

