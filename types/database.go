package types

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
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

func (db *DBInstance) AddTask(task Task) (int64, error) {
	query := "INSERT INTO tasks (name, time) VALUES (?, ?)"
	result, err := db.Instance.Exec(query, task.Name, task.Time.Format(time.DateTime))
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *DBInstance) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := db.Instance.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
