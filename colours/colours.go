package main

import (
	"fmt"
)

const (
	colourFmt = "\033[1;%dm%s\033[0m"

	openTagFmt = "\033[0;%dm"
	openTagBase = "\033["
	closeTag = "\033[0m"
)

type Code int

const (
	Black Code = iota + 30
	Red
	Green
)

type Attribute int

const (
	Reset Attribute = iota + 0
	Bold
)

type Colour struct {
	startTag string
}

func New(code Code, attributes... Attribute) Colour {
	attrStr := ""
	for _, attr := range attributes {
		attrStr += fmt.Sprintf("%d;", attr)
	}

	return Colour{
		startTag: fmt.Sprintf("%s%s%dm", openTagBase, attrStr, code),
	}
}

func (c Colour) Sprintf(msg string, args ...interface{}) string {
	return fmt.Sprintf("%s%s%s", c.startTag, fmt.Sprintf(msg, args...), closeTag)
}

func (c Colour) Printf(msg string, args... interface{}) {
	fmt.Printf("%s%s%s", c.startTag, fmt.Sprintf(msg, args...), closeTag)
}

func (c Colour) Highlight(args ...interface{}) []interface{} {
	highlighted := make([]interface{}, 0)

	for _, arg := range args {
		highlighted = append(highlighted, c.Sprintf("%s", arg))
	}

	return highlighted
}

func main() {
	red := New(Red, Bold)
	fmt.Println(red.Sprintf("This is red and bold"))

	green := New(Green)
	green.Printf("This is Green\n")

	args := green.Highlight("a", "b", "c")
	fmt.Printf("Here is some highlighted text %s %s %s this text is normal colour again", args...)
}
