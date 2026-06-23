package commands

import (
	"cynthia/service/database"
)

var db database.AppDatabase

func Init(d database.AppDatabase) {
	db = d
}
