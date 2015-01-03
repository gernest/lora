package logs

import (
	"bytes"
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/agtorre/gocolorize"
)

const (
	Red     = "red"
	Green   = "green"
	Black   = "black"
	Blue    = "blue"
	Magenta = "magenta"
	Cyan    = "cyan"

	INFO  = "INFO"
	TRAC  = "TRAC"
	ERRO  = "ERRO"
	WARN  = "WARN"
	SUCC  = "SUCC"
	DEBUG = "DEB"
	EVENT = "EVE"
	CRIT  = "CRIT"
)

func NewLoraLog() LoraLog {
	lola := LoraLog{}
	return lola
}

type LoraLog struct {
}

func (l *LoraLog) Debug(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	logThis(DEBUG, s)
}
func (l *LoraLog) Critical(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	logThis(CRIT, s)
}
func (l *LoraLog) Warning(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	logThis(INFO, s)
}
func (l *LoraLog) Info(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	logThis(INFO, s)
}
func (l *LoraLog) Success(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	logThis(SUCC, s)
}

func (l *LoraLog) Event(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	logThis(EVENT, s)
}
func colorize(s, color string) string {
	c := gocolorize.NewColor(color)
	return c.Paint(s)
}
func getFuncCall(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		_, filename := path.Split(file)
		s := fmt.Sprintf("[%s:%d] ", filename, line)
		return s
	}
	return ""
}
func getLoggingTime() string {
	t := time.Now()
	return t.Format("2006/01/02 15:04:05 ")
}
func logThis(level string, format string) {
	var colorLevel, callDepth, sep string
	switch level {
	case INFO:
		colorLevel = Green
	case DEBUG:
		colorLevel = Red
	case CRIT:
		colorLevel = Magenta
	case SUCC:
		colorLevel = Cyan
	case EVENT:
		colorLevel = Black
	case WARN:
		colorLevel = Blue
	case ERRO:
		colorLevel = Red
	}
	buf := new(bytes.Buffer)
	sep = "::"
	callDepth = getFuncCall(3)
	buf.WriteString(getLoggingTime())
	buf.WriteString(colorize("["+level+"]", colorLevel))
	buf.WriteString(sep)
	buf.WriteString(callDepth)
	buf.WriteString(colorize(format, colorLevel))
	fmt.Println(buf.String())
}
