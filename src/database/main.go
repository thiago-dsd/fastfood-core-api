package database

import (
	database_connect "github.com/thiago-dsd/fastfood-core-api/src/database/connect"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	return database_connect.DB
}
