package database

import (
	log "github.com/sirupsen/logrus"
	"privy/interfaces"
)

var DBConn interfaces.SQL

func init() {
	DBConn = NewSQL(GetConfig())
	if err := DBConn.MySQL(); err != nil {
		log.Error(err)
		return
	}
}
