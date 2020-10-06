package log

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

const (
	prefixFmt = "[%s]"
	infoEmoji = ":rainbow:"
	warnEmoji = ":rotating_light:"
	errEmoji  = ":fire:"
)

var (
	Name string

	styleInfo = newStyle(color.FgGreen)
	styleWarn = newStyle(color.FgYellow)
	styleErr  = newStyle(color.FgHiRed)

	writer io.Writer = os.Stdout

	tabWriter = tabwriter.NewWriter(writer, 0, 0, 1, ' ', tabwriter.AlignRight)
)

type style struct {
	bold   *color.Color
	italic *color.Color
	plain  *color.Color
}

func newStyle(c color.Attribute) style {
	return style{
		bold:   color.New(c, color.Bold),
		italic: color.New(c, color.Italic),
		plain:  color.New(c),
	}
}

func Init(name string) {
	Name = name
}

func Info(msg string, args ...interface{}) {
	write(styleInfo, infoEmoji, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	write(styleWarn, warnEmoji, msg, args...)
}

func Err(msg string, args ...interface{}) {
	write(styleErr, errEmoji, msg, args...)
}

func write(s style, symbol, msg string, args ...interface{}) {
	prefix := s.bold.Sprintf(prefixFmt, Name)
	timestamp := s.plain.Sprintf("%s", time.Now().Format(time.RFC3339))

	msg = emoji.Sprintf(fmt.Sprintf(msg, highlightArgs(s, args...)...))
	symbol =  emoji.Sprintf(symbol)

	fmt.Fprintln(tabWriter, fmt.Sprintf("%s\t%s\t%s \t%s", prefix, symbol, timestamp, msg))
	tabWriter.Flush()
}

func highlightArgs(s style, args ...interface{}) []interface{} {
	highlighted := make([]interface{}, 0)

	for _, arg := range args {
		highlighted = append(highlighted, s.italic.Sprintf("%s", arg))
	}

	return highlighted
}
