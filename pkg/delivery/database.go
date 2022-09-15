package delivery

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func SetupDatabase() (*sql.DB, error) {
	var (
		host     = viper.GetString("DATABASE_HOST")
		port     = viper.GetInt("DATABASE_PORT")
		user     = viper.GetString("DATABASE_USERNAME")
		password = viper.GetString("DATABASE_PASSWORD")
		db       = viper.GetString("DATABASE_NAME")
	)

	var err error
	var dbConnect = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)

	log.Printf("Connecting to dbName: %s@%s", db, host)
	open, err := sql.Open("postgres", dbConnect)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if err = open.Ping(); err != nil {
		return nil, err
	}
	return open, nil
}
