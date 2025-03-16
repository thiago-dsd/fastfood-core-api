package database_migrate

import (
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
	"github.com/pterm/pterm"
	"github.com/thiago-dsd/fastfood-core-api/src/database"
	_ "github.com/thiago-dsd/fastfood-core-api/src/database/migrations"
	order_entity "github.com/thiago-dsd/fastfood-core-api/src/order/entity"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

func Migrations() {
	// gooseBeforeAutomaticMigrations()
	automaticMigrations()
	gooseMigrations()
}

// Configures automatic migrations with ORM.
func automaticMigrations() {
	pterm.DefaultLogger.Info("Adding automatic migrations")
	err := database.Connection().AutoMigrate(
		&user_entity.User{},
		&order_entity.Order{},
	)
	if err != nil {
		pterm.DefaultLogger.Error(fmt.Sprintf("Unable to add automatic migrations: %s", err))
		os.Exit(1)
	}
	pterm.DefaultLogger.Info("Automatic migrations done")
}

// Executes goose migrations.
func gooseMigrations() {
	pterm.DefaultLogger.Info("Executing goose migrations...")
	// Configure Goose
	goose.SetDialect("postgres") // Set the database dialect

	// Run the migrations
	db, _ := database.Connection().DB()
	if err := goose.Up(db, "src/database/migrations"); err != nil {
		pterm.DefaultLogger.Error(fmt.Sprintf("Unable to execute goose migrations: %s", err))
		os.Exit(1)
	}

	pterm.DefaultLogger.Info("Goose migrations executed")
}

// Executes goose migrations.
// func gooseBeforeAutomaticMigrations() {
// 	pterm.DefaultLogger.Info("Executing goose before automatic migrations...")
// 	// Configure Goose
// 	goose.SetDialect("postgres") // Set the database dialect
//
// 	// Run the migrations
// 	db, _ := database.Connection().DB()
// 	if err := goose.Up(db, "src/database/migrations-before"); err != nil {
// 		pterm.DefaultLogger.Error(fmt.Sprintf("Unable to execute goose migrations before automatic: %s", err))
// 		os.Exit(1)
// 	}
//
// 	pterm.DefaultLogger.Info("Goose migrations before automatic executed")
// }
