package config

import (
	"github.com/thiago-dsd/fastfood-core-api/src/config/env"
	database_main "github.com/thiago-dsd/fastfood-core-api/src/database/main"
)

func Load() {
	env.Load()
	database_main.Database()
}
