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
	namespace string    = "funky-log"
	timeFmt   string    = time.RFC3339
	outW      io.Writer = os.Stdout

	tw = tabwriter.NewWriter(outW, 0, 0, 1, ' ', tabwriter.AlignRight)

	infoStyle = NewStyle(color.FgGreen, emoji.Sprintf(":white_check_mark:"))
	warnStyle = NewStyle(color.FgYellow, emoji.Sprintf(":warning: "))
	errStyle  = NewStyle(color.FgHiRed, emoji.Sprintf(":fire:"))
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
