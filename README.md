# funky-log
Simple logger for Go apps that doesn't take it's self too seriously.
- Support for emoji's
- Customize output text colours per log event level.
- Set the time layout.

## Sample output

![Alt text](preview.png?raw=true "Optional Title")

### Prerequisites
- Go version >= 1.15.2

### Getting started
```
go get github.com/daiLlew/funky-log
```

### Usage
Using the default setup:
```go
import "github.com/daiLlew/cli-fmt/log"

...

// Set the log namespace
log.Init("my-app")

// Log an info message with some emojis.
log.Info("time for :beer:and :pizza:")

// Log a warning
log.Warn("something is not quite right")

// Log and error with arguments.
log.Err("this is an error! %+v", errors.New("encountered an unexpected error"))
```
Create your own styles and customize the output:

```go
// Create a configuration to customize how the output should be formatted.
cfg := log.Configuration{
    Namespace: "my-app",
    TimeFmt:   time.RFC822,
    InfoStyle: log.NewStyle(color.FgHiCyan, ":unicorn_face:"),
    WarnStyle: log.NewStyle(color.FgHiBlue, ":tiger:"),
    ErrStyle:  log.NewStyle(color.FgHiMagenta, ":comet: "),
}

log.Customise(cfg)

log.Info("time for :beer:and :pizza:")
log.Warn("something is not quite right")
log.Err("this is an error! %+v", errors.New("encountered an unexpected error"))
```  
