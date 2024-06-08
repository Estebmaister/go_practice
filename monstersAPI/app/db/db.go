package db

import (
	"fmt"
	"log"

	"battle-of-monsters/app/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var CONN *gorm.DB

func Connect() *gorm.DB {
	if CONN != nil {
		return CONN
	}

	log.Println("Database connection started")

	var err error
	if CONN, err = gorm.Open(sqlite.Open(getDataBaseName()), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("failed to open the database connection. %w", err))
	}

	registerDBValidationsHooks(CONN)

	return CONN
}

func getDataBaseName() string {
	dn := config.ENV.DBName
	if dn == "" {
		log.Fatalln("database name is not defined to open a connection")
	}

	return dn
}
