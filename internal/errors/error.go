package errors

import (
	"github.com/jwalton/gchalk"
	"github.com/sirupsen/logrus"
	"privy/internal/options"
	"strconv"
)

func ErrorMessage(statusCode int, msg string) {
	caller := options.GetCaller(2)
	logrus.Errorf("[%s][%s][%s] %s",
		gchalk.BrightYellow(strconv.Itoa(statusCode)), gchalk.BrightWhite(caller.Fname),
		gchalk.BrightCyan(strconv.Itoa(caller.Line)),
		msg)
}
