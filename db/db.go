package db

import (
	"database/sql"
	"event-management/utils/env"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Connection *sql.DB
	err        error
)

func Setup() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable timezone=%s",
		env.Get("DB_HOST", ""),
		env.Get("DB_PORT", ""),
		env.Get("DB_USER", ""),
		env.Get("DB_SECRET", ""),
		env.Get("DB_NAME", ""),
		env.Get("LOG_TIMEZONE", ""),
	)

	Connection, err = sql.Open(env.Get("DB_DRIVER", ""), psqlInfo)
	if err != nil {
		panic(err)
	}

	err = Connection.Ping()
	if err != nil {
		panic(err)
	}
}
