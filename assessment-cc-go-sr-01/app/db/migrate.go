package db

import (
	"database/sql"
	"log"

	"battle-of-monsters/app/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Source type used
	_ "modernc.org/sqlite"
)

func Up() {
	log.Println("migration up started")

	var migration = buildMigrate()

	migrationError := migration.Up()

	if migrationError != nil {
		log.Fatalln("fail when migration up execution : ", migrationError.Error())
	}

	log.Println("migration up finished")
}

func Down() {
	log.Println("migration down started")

	var migration = buildMigrate()

	migrationError := migration.Down()

	if migrationError != nil {
		log.Fatalln("fail when migration down execution : ", migrationError.Error())
	}

	log.Println("migration down finished")
}

func buildMigrate() *migrate.Migrate {
	config.Load()

	var driver database.Driver

	db, err := sql.Open(config.ENV.DBDriver, config.ENV.DBName)
	if err != nil {
		log.Fatalln("error opening database connection during the migration process")
	}

	if driver, err = sqlite.WithInstance(db, &sqlite.Config{}); err != nil {
		log.Fatalln("error getting instance driver to migrations : ", err.Error())
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://./app/db/migrations",
		config.ENV.DBDriver, driver)
	if err != nil {
		log.Fatalln("error creating new migrate instance : ", err.Error())
	}

	return migration
}
