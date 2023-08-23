package driver

import (
	"database/sql"
	"os"

	"github.com/pkg/errors"

	"github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

// Регистрирует соединение с БД
func New() (*MySQL, error) {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to mysql")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping db")
	}

	return &MySQL{
		DB: db,
	}, nil
}

func (db *MySQL) Close() {
	db.DB.Close()
}
