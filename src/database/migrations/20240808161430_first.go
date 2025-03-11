package migrations

import (
	"context"
	"database/sql"

	"github.com/cogniia/core-api-template/src/database"
	entity "github.com/cogniia/core-api-template/src/user/entity"
	"github.com/pressly/goose/v3"
	"github.com/pterm/pterm"
)

func init() {
	goose.AddMigrationContext(upFirst, downFirst)
}

func upFirst(ctx context.Context, tx *sql.Tx) error {
	var existingUser entity.User
	err := database.Connection().First(&existingUser).Error

	if err != nil {
		pterm.DefaultLogger.Info("No user in database.")
	} else {
		pterm.DefaultLogger.Info("User found in database.")
	}

	return nil
}

func downFirst(ctx context.Context, tx *sql.Tx) error {
	var existingUser entity.User
	err := database.Connection().First(&existingUser).Error

	if err != nil {
		pterm.DefaultLogger.Info("No user in database.")
	} else {
		pterm.DefaultLogger.Info("User found in database.")
	}

	return nil
}
