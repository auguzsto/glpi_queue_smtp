package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofor-little/env"
)

func Con() *sql.DB {
	DBROOT := env.Get("DBROOT", "DEFAULT_VALUE")
	DBHOST := env.Get("DBHOST", "DEFAULT_VALUE")
	DBPORT := env.Get("DBPORT", "DEFAULT_VALUE")
	DBPASS := env.Get("DBPASS", "DEFAULT_VALUE")
	DBGLPI := env.Get("DBGLPI", "DEFAULT_VALUE")

	dataSourceName := DBROOT + ":" + DBPASS + "@tcp(" + DBHOST + ":" + DBPORT + ")/" + DBGLPI
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
