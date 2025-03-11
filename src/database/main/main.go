package database_main

import (
	database_connect "github.com/cogniia/core-api-template/src/database/connect"
	database_migrate "github.com/cogniia/core-api-template/src/database/migrate"
)

func Database() {
	database_connect.Connect()
	database_migrate.Migrations()
}
