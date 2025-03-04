package database

import (
	"context"
	"sync"

	"github.com/LuanTenorio/learn-api/internal/config"
	"github.com/LuanTenorio/learn-api/internal/database/sqlc"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type pgDatabase struct {
	Db      *pgx.Conn
	Queries *sqlc.Queries
}

var (
	once           sync.Once
	dbSqlcInstance *pgDatabase
)

func NewPGDatabase(conf *config.Config) Database {
	ctx := context.Background()
	once.Do(func() {
		db, err := pgx.Connect(ctx, conf.Db.Url)
		if err != nil {
			panic("failed to connect database")
		}

		dbSqlcInstance = &pgDatabase{Db: db, Queries: sqlc.New(db)}
	})

	return dbSqlcInstance
}

func (p *pgDatabase) GetDb() *pgx.Conn {
	return dbSqlcInstance.Db
}

func (p *pgDatabase) GetQueries() *sqlc.Queries {
	return dbSqlcInstance.Queries
}
