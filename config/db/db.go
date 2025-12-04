package db

import (
	"GoTwitter/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()

	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.DBName = env.GetString("DB_NAME", "dev_db")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Passwd = env.GetString("DB_PASSWORD", "Maclocal@123")
	cfg.User = env.GetString("DB_USER", "root")

	fmt.Println("connecting to db", cfg.FormatDSN())
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
		return nil, err
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Error pinging to database")
		return nil, pingErr
	}

	fmt.Println("connected to database successfully")

	return db, nil
}
