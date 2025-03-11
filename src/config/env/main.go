package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

func Load() {
	loadEnv()
	loadAuthEnv()
	loadDbEnv()
	loadServerEnv()
}

func loadEnv() {
	pterm.DefaultLogger.Info(
		"Loading production environment file...",
	)
	err := godotenv.Load(".env")
	if err != nil {
		pterm.DefaultLogger.Warn(
			"No production environment found. Please create a .env file in the root of the project if you want a production context",
		)
		pterm.DefaultLogger.Info(
			"Loading development environment file...",
		)
		err = godotenv.Load(".env.local")
		if err != nil {
			pterm.DefaultLogger.Error(
				fmt.Sprintf("Some error occurred loading the environment file: %s", err),
			)
			os.Exit(1)
		}
	}

	pterm.DefaultLogger.Info(
		"Environment file successfully set",
	)
}
