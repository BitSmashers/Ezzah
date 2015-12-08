package utils

import (
	"github.com/op/go-logging"
)

type Logger struct {
	log *logging.Logger
}

var LOG = InitLog()

// Ezzah logging general format
var format = logging.MustStringFormatter(
	"%{id:03x} %{shortfile} %{level:.6s} ▶ %{color:reset} %{message}",
)

func InitLog() *logging.Logger {
	logging.SetFormatter(format)
	var log, err = logging.GetLogger("Ezzah")

	//Bypass HandleError(err) to avoid init loop
	if err != nil {
		panic(err)
	}
	log.Debug("Logger successfuly init")
	return log
}
// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}