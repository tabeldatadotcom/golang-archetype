package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"tabeldatadotcom/archetype/backend-api/config"
)

func main() {
	var (
		dbEnv       = viper.GetString("DATABASE_NAME")
	)

	config.SetupEnvironment()

	db, dbErr := config.SetupDatabase()
	defer db.Close()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	driver, driverErr := postgres.WithInstance(db.DB, &postgres.Config{})
	if driverErr != nil {

	}
	m, errMigrate := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		dbEnv,
		driver,
	)
	if errMigrate != nil {
		log.Fatal(errMigrate)
	}

	_ = m.Up()
	log.Println("---do-migrate---")

	version, dirty, verErr := m.Version()
	if verErr != nil {
		log.Fatal(verErr)
	}
	log.Printf("version: %d, dirty: %t", version, dirty)

}
