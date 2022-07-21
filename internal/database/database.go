package database

import (
	"github.com/fajarardiyanto/boiler-monolith-rest-api/interfaces"
	log "github.com/sirupsen/logrus"
)

var DBConn interfaces.SQL

func init() {
	DBConn = NewSQL(GetConfig())
	if err := DBConn.MySQL(); err != nil {
		log.Error(err)
		return
	}
}
