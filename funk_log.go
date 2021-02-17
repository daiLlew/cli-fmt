package funkylog

import (
	"fmt"
	"io"
	"os"
	"strings"
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
	namespace string    = "funky-log"
	timeFmt   string    = time.RFC3339
	outW      io.Writer = os.Stdout

	tw = tabwriter.NewWriter(outW, 0, 0, 1, ' ', tabwriter.AlignRight)

	infoStyle = NewStyle(color.FgGreen, emoji.Sprintf(":white_check_mark:"))
	warnStyle = NewStyle(color.FgYellow, emoji.Sprintf(":warning: "))
	errStyle  = NewStyle(color.FgHiRed, emoji.Sprintf(":fire:"))

	formats = []string{
		"%v",
		"%#v",
		"%T",
		"%t",
		"%b",
		"%c",
		"%d",
		"%o",
		"%0",
		"%q",
		"%x",
		"%X",
		"%U",
		"%b",
		"%e",
		"%E",
		"%f",
		"%F",
		"%g",
		"%G",
		"%s",
		"%p",
	}
)

type Configuration struct {
	Namespace string
	TimeFmt   string
	InfoStyle Style
	WarnStyle Style
	ErrStyle  Style
}

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
	namespace = name
}

func Customise(cfg Configuration) {
	if cfg.Namespace != "" {
		namespace = cfg.Namespace
	}

	if cfg.TimeFmt != "" {
		timeFmt = cfg.TimeFmt
	}

	infoStyle = cfg.InfoStyle
	warnStyle = cfg.WarnStyle
	errStyle = cfg.ErrStyle
}

func Info(msg string, args ...interface{}) {
	infoStyle.Write(tw, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	warnStyle.Write(tw, msg, args...)
}

func Err(msg string, args ...interface{}) {
	errStyle.Write(tw, msg, args...)
}

func (s Style) Sprintf(msg string, args ...interface{}) string {
	prefix := s.Bold.Sprintf(prefixFmt, namespace)

	ts := s.Plain.Sprintf("%s", time.Now().Format(timeFmt))

	msg = s.highlightArgs(msg, args...)

	msg = emoji.Sprintf(msg)

	return fmt.Sprintf(outputFmt, prefix, s.Emoji, ts, msg)
}

func (s Style) Write(tw *tabwriter.Writer, msg string, args ...interface{}) {
	fmt.Fprintln(tw, s.Sprintf(msg, args...))
	tw.Flush()
}

// If the msg string contains any of the supported Golang formats then wrap it in an open/close colour tag.
// Example:
// If message contains "%s" replace it with "[start_colour]%s[end_colour]"
// This allows the argument to be formatted as desired using the standard fmt library but the inclusion of the
// start/end colour tags will "highlight" the string value when its output
func (s Style) highlightArgs(msg string, args ...interface{}) string {
	for _, f := range formats {
		if strings.Contains(msg, f) {

			msg = strings.ReplaceAll(msg, f, s.Italic.Sprintf("%s", f))
		}
	}

	return fmt.Sprintf(msg, args...)
}
