package types

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

type DBInstance struct {
	Instance *sql.DB
}

func NewDBInstance() *sql.DB {
	log.Print(viper.GetString("db"))
	sqlDB, err := sql.Open("sqlite", viper.GetString("db"))
	if err != nil {
		log.Print(err)
	}

	dbinit, err := os.ReadFile("../configs/sql_init.sql")
	if err != nil {
		log.Print(err)
	}
	if _, err = sqlDB.Exec(string(dbinit)); err != nil {
		log.Print(err)
	}
	log.Print("Successfully connxted to DB")
	return sqlDB

}
