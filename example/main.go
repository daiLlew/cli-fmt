package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/daiLlew/cli-fmt/log"
	"github.com/fatih/color"
)

func main() {
	fmt.Println()
	fmt.Println("default configuration")

	log.Init("my-app")
	log.Info("time for :beer:and :pizza:")
	log.Warn("something is not quite right")
	log.Err("this is an error! %+v", errors.New("encountered an unexpected error"))

	fmt.Println()
	fmt.Println("customized configuration")

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

	fmt.Println()
}
