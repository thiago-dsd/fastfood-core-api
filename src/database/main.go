package database

import (
	database_connect "github.com/cogniia/core-api-template/src/database/connect"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	return database_connect.DB
}
