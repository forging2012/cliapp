package color

import (
	"fmt"
	"strings"
)

// value is a defined style name
type Tag string

// Print
func (tg Tag) Print(args ...interface{}) {
	str := buildColoredText(GetStyleCode(string(tg)), args...)
	fmt.Print(str)
}

// Printf
func (tg Tag) Printf(format string, args ...interface{}) {
	str := buildColoredText(GetStyleCode(string(tg)), fmt.Sprintf(format, args...))
	fmt.Print(str)
}

// Println
func (tg Tag) Println(args ...interface{}) {
	str := buildColoredText(GetStyleCode(string(tg)), args...)
	fmt.Println(str)
}

// Tips will add color for all text
// value is a defined style name
type Tips string

// Print
func (t Tips) Print(args ...interface{}) {
	tag := string(t)
	str := buildColoredText(
		GetStyleCode(tag),
		strings.ToUpper(tag), ": ", fmt.Sprint(args...),
	)

	fmt.Println(str)
}

// Println
func (t Tips) Println(args ...interface{}) {
	t.Print(args...)
}

// Printf
func (t Tips) Printf(format string, args ...interface{}) {
	tag := string(t)
	str := buildColoredText(
		GetStyleCode(tag),
		strings.ToUpper(tag), ": ", fmt.Sprintf(format, args...),
	)

	fmt.Println(str)
}

// LiteTips will only add color for tag name
// value is a defined style name
type LiteTips string

// Print
func (t LiteTips) Print(args ...interface{}) {
	tag := string(t)
	str := buildColoredText(GetStyleCode(tag), strings.ToUpper(tag), ":")

	fmt.Println(str, fmt.Sprint(args...))
}

// Println
func (t LiteTips) Println(args ...interface{}) {
	t.Print(args...)
}

// Printf
func (t LiteTips) Printf(format string, args ...interface{}) {
	tag := string(t)
	str := buildColoredText(GetStyleCode(tag), strings.ToUpper(tag), ":")

	fmt.Println(str, fmt.Sprintf(format, args...))
}

// Logger console logger
type Logger struct {
	style  string
	fields map[string]string
}

// log level to cli color style
var LogLevel2tag = map[string]string{
	"info":    "info",
	"warn":    "warning",
	"warning": "warning",
	"debug":   "cyan",
	"notice":  "notice",
	"error":   "error",
}

func NewLog(fields map[string]string) *Logger {
	return &Logger{Info, fields}
}

func (l *Logger) Info(args ...interface{}) {

}

func (l *Logger) Log(args ...interface{}) {

}
