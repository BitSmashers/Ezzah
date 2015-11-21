package utils

import (
	"github.com/op/go-logging"
)

var LOG = InitLog()

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func InitLog() *logging.Logger {
	var log, err = logging.GetLogger("Ezzah")
	HandleError(err)
	log.Debug("Logger successfuly init")
	return log
}
// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}