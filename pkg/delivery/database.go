package delivery

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func SetupDatabase() error {
	var err error
	var dbConnect = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", viper.GetString("DATABASE_HOST"), viper.GetInt("DATABASE_PORT"), viper.GetString("DATABASE_USERNAME"), viper.GetString("DATABASE_PASSWORD"), viper.GetString("DATABASE_NAME"))

	log.Printf("Connecting to dbName: %s@%s", viper.GetString("DATABASE_NAME"), viper.GetString("DATABASE_HOST"))
	open, err := sql.Open("postgres", dbConnect)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	if err = open.Ping(); err != nil {
		return err
	}
	return nil
}
