package logs

import (
	"bytes"
	"fmt"
	"path"
	"runtime"

	"github.com/agtorre/gocolorize"
)

func NewLoraLog() LoraLog {
	lola := LoraLog{}
	lola.debug = "[ERRO]"
	lola.info = "[INFO]"
	lola.critical = "[TRAC]"
	lola.warning = "[WARN"
	return lola
}

type LoraLog struct {
	info     string
	debug    string
	critical string
	warning  string
}

func (l *LoraLog) Debug(format string, a ...interface{}) {
	buf := new(bytes.Buffer)
	buf.WriteString(l.debug)
	buf.WriteString(" ")
	buf.WriteString(getFuncCall(2))

	buf.WriteString(colorize(fmt.Sprintf(format, a...), "red"))
	buf.WriteString("\n")
	ColorLog(buf.String())
}
func (l *LoraLog) Critical(format string, a ...interface{}) {
	buf := new(bytes.Buffer)
	buf.WriteString(l.critical)
	buf.WriteString(" ")
	buf.WriteString(getFuncCall(2))

	buf.WriteString(colorize(fmt.Sprintf(format, a...), "black"))
	buf.WriteString("\n")
	ColorLog(buf.String())
}
func (l *LoraLog) Warning(format string, a ...interface{}) {
	buf := new(bytes.Buffer)
	buf.WriteString(l.debug)
	buf.WriteString(" ")
	buf.WriteString(getFuncCall(2))

	buf.WriteString(colorize(fmt.Sprintf(format, a...), "blue"))
	buf.WriteString("\n")
	ColorLog(buf.String())
}
func (l *LoraLog) Info(format string, a ...interface{}) {
	buf := new(bytes.Buffer)
	buf.WriteString(l.info)
	buf.WriteString(" ")
	buf.WriteString(getFuncCall(2))
	buf.WriteString(colorize(fmt.Sprintf(format, a...), "green"))
	buf.WriteString("\n")
	ColorLog(buf.String())
}
func (l *LoraLog) Success(format string, a ...interface{}) {
	buf := new(bytes.Buffer)
	buf.WriteString("[SUCC]")
	buf.WriteString(" ")
	buf.WriteString(colorize(fmt.Sprintf(format, a...), "cyan"))
	buf.WriteString("\n")
	ColorLog(buf.String())
}

func (l *LoraLog) Event(format string, a ...interface{}) {
	buf := new(bytes.Buffer)
	buf.WriteString("[EVEN]")
	buf.WriteString(" ")
	buf.WriteString(colorize(fmt.Sprintf(format, a...), "black"))
	buf.WriteString("\n")
	ColorLog(buf.String())
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
