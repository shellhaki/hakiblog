package config

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Conndb() error {
	cstring := os.Getenv("POSTGRES_URI")
	if cstring == "" {
		return errors.New("Connection string missing")
	}
	db, err := sql.Open("pgx", cstring)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return err
	}
	log.Println("database connected successfully")
	DB = db
	return nil
}
