package database

import (
	"sync"

	"github.com/LuanTenorio/learn-api/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresDatabase struct {
	Db *sqlx.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	once.Do(func() {
		db, err := sqlx.Connect("postgres", conf.Db.Url)
		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = &postgresDatabase{Db: db}
	})

	return dbInstance
}

func (p *postgresDatabase) GetDb() *sqlx.DB {
	return dbInstance.Db
}
