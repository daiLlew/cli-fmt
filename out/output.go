package out

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

const (
	prefixFmt = "[%s]"
	tickEmoji = ":white_check_mark:"
	warnEmoji = ":warning:"
	fireEmoji = ":fire:"
)

var (
	Name        string
	infoColour  = color.New(color.FgGreen, color.Bold)
	warnColour  = color.New(color.FgYellow, color.Bold)
	errorColour = color.New(color.FgHiRed, color.Bold)

	writer io.Writer = os.Stdout
)

func Init(name string) {
	Name = name
}

func Info(msg string, args ...interface{}) {
	write(infoColour, tickEmoji, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	write(warnColour, warnEmoji, msg, args...)
}

func Err(msg string, args ...interface{}) {
	write(errorColour, fireEmoji, msg, args...)
}

func write(c *color.Color, symbol, msg string, args ...interface{}) {
	prefix := c.Sprintf(prefixFmt, Name)
	msg = fmt.Sprintf(msg, highlightArgs(c, args...)...)
	msg = emoji.Sprintf("%s %s %s", prefix, symbol, msg)

	fmt.Fprintln(writer, msg)
}

func highlightArgs(c *color.Color, args ...interface{}) []interface{} {
	highlighted := make([]interface{}, 0)
	for _, arg := range args {
		highlighted = append(highlighted, c.Sprintf("%s", arg))
	}

	return highlighted
}
