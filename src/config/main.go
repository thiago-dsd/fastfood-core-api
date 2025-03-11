package config

import (
	"github.com/cogniia/core-api-template/src/config/env"
	database_main "github.com/cogniia/core-api-template/src/database/main"
)

func Load() {
	env.Load()
	database_main.Database()
}
