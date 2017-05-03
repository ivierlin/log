package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var start = time.Now()

func Log(prefix, s string, a ...interface{}) {
	fmt.Printf("%6.6f %6s\t%-12s%s", time.Since(start).Seconds(), prefix, s, fmt.Sprintln(a...))
}

var (
	Info  = func(s string, a ...interface{}) { Log(color.BlueString("INFO"), s, a...) }
	Debug = func(s string, a ...interface{}) { Log(color.GreenString("DEBUG"), s, a...) }
	Warn  = func(s string, a ...interface{}) { Log(color.YellowString("WARNING"), s, a...) }
	Err   = func(s string, a ...interface{}) { Log(color.RedString("ERROR"), s, a...) }
)
