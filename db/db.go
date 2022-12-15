package db

import (
	"database/sql"
	"event-management/utils/env"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
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

func Migrate() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(Connection, env.Get("DB_DRIVER", "postgres"), migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	fmt.Println("Applied", n, " migrations!")
}
