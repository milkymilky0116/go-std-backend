package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

var lock = &sync.Mutex{}
var DBQueries *Queries
var ctx context.Context

func InitDB() (*Queries, context.Context) {

	if DBQueries == nil {
		lock.Lock()
		defer lock.Unlock()
		if DBQueries == nil {
			log.Println("Initializing DB Table.")
			ctx = context.Background()

			db, err := sql.Open("sqlite3", ":memory:")

			if err != nil {
				log.Fatal(err)
			}

			if _, err := db.ExecContext(ctx, ddl); err != nil {
				log.Fatal(err)
			}

			DBQueries = New(db)

		} else {
			log.Println("DB already initialized.")
		}
	} else {
		log.Println("DB already initialized.")
	}

	return DBQueries, ctx
}
