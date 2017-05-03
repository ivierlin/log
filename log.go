package log

import (
	"fmt"
	"time"
)

const (
	Black = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

var start = time.Now()

func Log(prefix string, a ...interface{}) {
	fmt.Printf("%6.6f  %-18s%s", time.Since(start).Seconds(), prefix, fmt.Sprintln(a...))
}

func Prefix(s string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
}

var (
	Info  = func(a ...interface{}) { Log(Prefix("INFO", Blue), a...) }
	Debug = func(a ...interface{}) { Log(Prefix("DEBUG", Green), a...) }
	Warn  = func(a ...interface{}) { Log(Prefix("WARNING", Yellow), a...) }
	Err   = func(a ...interface{}) { Log(Prefix("ERROR", Red), a...) }
)
