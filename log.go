package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var red = color.New(color.FgHiRed, color.Bold).SprintFunc()
var yellow = color.New(color.FgHiYellow).SprintFunc()

func log(t string, s string) {
	n := time.Now().Format("2006-01-02T15:04:05-0700")
	t = "[" + t + "]"

	if t == "[FATAL]" {
		t = red(t)
		s = red(s)
	} else if t == "[WARN]" {
		t = yellow(t)
		s = yellow(s)
	}

	fmt.Printf("%s: %s %s\n", n, t, s)
}
