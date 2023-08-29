package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func OpenEnv() error {
	err := godotenv.Load(".env")
	return err
}

func CreateConnector() string {
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	return ConnectionString
}

func GetDB() (*sql.DB, error) {
	return sql.Open("mysql", CreateConnector())
}
