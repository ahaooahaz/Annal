package storage

import (
	"database/sql"
	"sync"

	"github.com/ahaooahaz/Annal/binaries/config"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

var (
	instance *sqlx.DB
	once     sync.Once
)

type DB interface {
	Select(interface{}, string, ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
	Get(interface{}, string, ...interface{}) error
}

func GetInstance() (dv *sqlx.DB) {
	once.Do(func() {
		var err error
		dv, err = sqlx.Open("sqlite3", config.DBPATH)
		if err != nil {
			panic(err.Error())
		}

		instance = dv
		sqlbuilder.DefaultFlavor = sqlbuilder.SQLite
	})

	dv = instance
	return
}
