package database_main

import (
	database_connect "github.com/thiago-dsd/fastfood-core-api/src/database/connect"
	database_migrate "github.com/thiago-dsd/fastfood-core-api/src/database/migrate"
)

func Database() {
	database_connect.Connect()
	database_migrate.Migrations()
}
