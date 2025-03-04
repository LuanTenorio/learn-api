package database

import (
	"github.com/LuanTenorio/learn-api/internal/database/sqlc"
	"github.com/jackc/pgx/v5"
)

type Database interface {
	GetDb() *pgx.Conn
	GetQueries() *sqlc.Queries
}

var (
	ErrNoRows = pgx.ErrNoRows
)
