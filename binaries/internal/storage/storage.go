package storage

import (
	"sync"

	"github.com/AHAOAHA/Annal/binaries/internal/config"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

var (
	instance *sqlx.DB
	once     sync.Once
)

func getInstance() (dv *sqlx.DB) {
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
