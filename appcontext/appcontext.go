package appcontext

import (
	"userland/config"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type appContext struct {
	db *sqlx.DB
}

var context *appContext

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func InitContext() {
	err := godotenv.Load()
	check(err)
	db, err := sqlx.Connect("postgres", config.GetDatabaseConnectionString())
	check(err)
	context = &appContext{db}
}

func GetDB() *sqlx.DB {
	return context.db
}
