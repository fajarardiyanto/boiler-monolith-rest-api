package parsing

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func StrToBoolEnv(s string) bool {
	boolValue, err := strconv.ParseBool(os.Getenv(s))
	if err != nil {
		logrus.Error(err)
	}
	return boolValue
}

func StrToIntEnv(s string) int {
	intValue, err := strconv.Atoi(os.Getenv(s))
	if err != nil {
		logrus.Error(err)
	}
	return intValue
}

func StrEnv(s string) string {
	return os.Getenv(s)
}
