package main

import (
	"errors"

	"github.com/daiLlew/cli-fmt/log"
)

func main() {
	log.Init("dp-cli")
	log.Info("time for :beer:and :pizza:")
	log.Warn("something is not quite right")
	log.Err("this is an error! %+v end.", errors.New("bork"))
}
