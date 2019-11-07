package appcontext

import (
	"github.com/jmoiron/sqlx"
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
	db, err := sqlx.Open("postgres", "dbname=userland host=localhost port=5432 sslmode=disable")
	check(err)
	err = db.Ping()
	check(err)
	context = &appContext{db}
}

func GetDB() *sqlx.DB {
	return context.db
}
