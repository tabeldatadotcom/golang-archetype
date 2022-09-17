package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func SetupDatabase() (*sqlx.DB, error) {
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
	connect, err := sqlx.Open("postgres", dbConnect)
	if err != nil {
		return nil, err
	}

	if err = connect.Ping(); err != nil {
		return nil, err
	}
	return connect, nil
}
