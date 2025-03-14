package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/pterm/pterm"
	"github.com/thiago-dsd/fastfood-core-api/src/database"
	entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
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
