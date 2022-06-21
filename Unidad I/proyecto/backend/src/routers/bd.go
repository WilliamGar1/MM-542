package routers

import (
	"database/sql"
	"proyecto/src/config"
)

var db *sql.DB
var err error

func InitDB() {
	db, err = config.GetDB()
}
