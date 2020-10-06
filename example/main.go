package main

import (
	"errors"

	"github.com/daiLlew/cli-fmt/out"
)

func main() {
	out.Init("dp-cli")
	out.Info("time for :beer:and :pizza:")
	out.Warn("something is not quite right")
	out.Err("this is an error! %+v end.", errors.New("bork"))
}
