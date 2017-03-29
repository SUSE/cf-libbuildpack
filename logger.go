package libbuildpack

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Logger interface {
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	BeginStep(format string, args ...interface{})
	Protip(tip string, help_url string)
	GetOutput() io.Writer
	SetOutput(w io.Writer)
}

type logger struct {
	w io.Writer
}

func NewLogger() Logger {
	return &logger{w: os.Stdout}
}

func (l *logger) Info(format string, args ...interface{}) {
	l.printWithHeader(none, "       ", format, args...)
}

func (l *logger) Warning(format string, args ...interface{}) {
	l.printWithHeader(yellow, "       **WARNING** ", format, args...)

}
func (l *logger) Error(format string, args ...interface{}) {
	l.printWithHeader(red, "       **ERROR** ", format, args...)
}

func (l *logger) BeginStep(format string, args ...interface{}) {
	l.printWithHeader(none, "-----> ", format, args...)
}

func (l *logger) Protip(tip string, helpURL string) {
	l.printWithHeader(none, "       PRO TIP: ", "%s", tip)
	l.printWithHeader(none, "       Visit ", "%s", helpURL)
}

func (l *logger) printWithHeader(color func(string) string, header string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	msg = strings.Replace(msg, "\n", "\n       ", -1)
	uncolored := fmt.Sprintf("%s%s\n", header, msg)
	fmt.Fprintf(l.w, color(uncolored))
}

func (l *logger) GetOutput() io.Writer {
	return l.w
}

func (l *logger) SetOutput(w io.Writer) {
	l.w = w
}

func red(uncolored string) string {
	return fmt.Sprintf("\033[31;1m%s\033[0m", uncolored)
}

func yellow(uncolored string) string {
	return fmt.Sprintf("\033[33;1m%s\033[0m", uncolored)
}

func none(uncolored string) string {
	return uncolored
}

var Log = &logger{w: os.Stdout}
