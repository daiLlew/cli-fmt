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
	outputFmt = "%s\t%s\t%s \t%s"
)

var (
	Name       string    = "funky-log"
	TimeLayout string    = time.RFC3339
	outW       io.Writer = os.Stdout

	tw        = tabwriter.NewWriter(outW, 0, 0, 1, ' ', tabwriter.AlignRight)

	styleInfo = NewStyle(color.FgGreen, emoji.Sprintf(":white_check_mark:"))
	styleWarn = NewStyle(color.FgYellow, emoji.Sprintf(":warning: "))
	styleErr  = NewStyle(color.FgHiRed, emoji.Sprintf(":fire:"))
)

type Style struct {
	Emoji  string
	Bold   *color.Color
	Italic *color.Color
	Plain  *color.Color
}

func NewStyle(c color.Attribute, emojiStr string) Style {
	return Style{
		Emoji:  emoji.Sprintf(emojiStr),
		Bold:   color.New(c, color.Bold),
		Italic: color.New(c, color.Italic),
		Plain:  color.New(c),
	}
}

func Init(name string) {
	Name = name
}

func Info(msg string, args ...interface{}) {
	styleInfo.Write(tw, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	styleWarn.Write(tw, msg, args...)
}

func Err(msg string, args ...interface{}) {
	styleErr.Write(tw, msg, args...)
}

func (s Style) Sprintf(msg string, args ...interface{}) string {
	prefix := s.Bold.Sprintf(prefixFmt, Name)

	ts := s.Plain.Sprintf("%s", time.Now().Format(TimeLayout))

	msg = emoji.Sprintf(fmt.Sprintf(msg, s.highlightArgs(args...)...))

	return fmt.Sprintf(outputFmt, prefix, s.Emoji, ts, msg)
}

func (s Style) Write(tw *tabwriter.Writer, msg string, args ...interface{}) {
	fmt.Fprintln(tw, s.Sprintf(msg, args...))
	tw.Flush()
}

func (s Style) highlightArgs(args ...interface{}) []interface{} {
	highlighted := make([]interface{}, 0)

	for _, arg := range args {
		highlighted = append(highlighted, s.Italic.Sprintf("%s", arg))
	}

	return highlighted
}
