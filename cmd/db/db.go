package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/milkymilky0116/go-std-backend/cmd/cli"
)

func DBinit(c cli.Config) (*sql.DB, error) {
	cfg := mysql.Config{
		User:      c.User,
		Passwd:    c.Password,
		Net:       "tcp",
		Addr:      fmt.Sprintf("127.0.0.1:%d", c.Port),
		DBName:    c.Dbname,
		ParseTime: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	return db, nil
}
