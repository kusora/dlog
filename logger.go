package dlog

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"io"
	"os"
	"github.com/kusora/raven-go"
)

const (
	FATAL = 0
	PANIC = 1
	ERROR = 2
	WARN = 3
	INFO = 4
	DEBUG = 5
)

var Level int64 = INFO
var SentryLevel int64 = ERROR

var hasSentry bool = false

var lg *Logger

func init() {
	lg = New(os.Stdout, "")
}

func SetSentry(DSN string, level int64) {
	if len(DSN) > 0 {
		raven.SetDSN(DSN)
		hasSentry = true
	}
	SentryLevel = level
}

func SetWriter(writer io.Writer) {
	lg.out = writer
}

func Info(format string, v ...interface{}) {
	if Level >= INFO {
		lg.Output(2, fmt.Sprintf("[INFO] " + format, v...))
	}

	if SentryLevel >= INFO && hasSentry {
		raven.CaptureMessage("[INFO] " + fmt.Sprintf(format, v...), nil)
	}
}

func InfoC(format string, v ...interface{}) {
	if Level >= INFO {
		lg.Output(2, fmt.Sprintf("[INFO] " + format, v...))
	}
	raven.CaptureMessage(fmt.Sprintf(format, v...), nil)
}

func Println(v ...interface{}) {
	lg.Output(2, fmt.Sprint(v...))
}

func Warn(format string, v ...interface{}) {
	if Level >= WARN {
		escapeCode := color.Colorize("y")
		io.WriteString(os.Stdout, escapeCode)
		line := color.Sprintf("[WARN] " + format, v...)
		lg.Output(2, line)
	}
	if SentryLevel >= WARN && hasSentry  {
		raven.CaptureMessage("[WARN] " + fmt.Sprintf(format, v...), nil)
	}
}

func Error(format string, v ...interface{}) {
	if Level >= ERROR {
		escapeCode := color.Colorize("r")
		io.WriteString(os.Stdout, escapeCode)
		line := color.Sprintf("[ERROR] " + format, v...)
		lg.Output(2, line)
	}

	if SentryLevel >= ERROR && hasSentry {
		raven.CaptureMessage("[ERROR] " + fmt.Sprintf(format, v...), nil)
	}
}

func ErrorN(n int, format string, v ...interface{}) {
	if Level >= ERROR {
		lg.Output(2 + n, fmt.Sprintf("[ERROR] " + format, v...))
	}

	if SentryLevel >= ERROR && hasSentry {
		raven.CaptureMessage("[ERROR] " + fmt.Sprintf(format, v...), nil)
	}
}

func Debug(format string, v ...interface{}) {
	if Level >= DEBUG {
		lg.Output(2, fmt.Sprintf("[DEBUG] " + format, v...))
	}

	if SentryLevel >= DEBUG && hasSentry {
		raven.CaptureMessage("[DEBUG] " + fmt.Sprintf(format, v...), nil)
	}
}

func Fatal(format string, v ...interface{}) {
	if Level >= FATAL {
		lg.Output(2, fmt.Sprintf("[FATAL] " + format, v...))
		os.Exit(1)
	}

	if SentryLevel >= FATAL && hasSentry {
		raven.CaptureMessage("[FATAL] " + fmt.Sprintf(format, v...), nil)
	}
}

func Fatalln(v ...interface{}) {
	if Level >= FATAL {
		lg.Output(2, fmt.Sprint(v...))
		os.Exit(1)
	}

	if SentryLevel >= FATAL && hasSentry {
		raven.CaptureMessage(fmt.Sprintf("fatal, %+v", v), nil)
	}
}

func Panic(format string, v ...interface{}) {
	if Level >= PANIC {
		s := fmt.Sprintf("[PANIC] " + format, v...)
		lg.Output(2, s)
		panic(s)
	}

	if SentryLevel >= PANIC && hasSentry {
		raven.CaptureMessage(fmt.Sprintf("[PANIC] " + format, v...), nil)
	}
}
