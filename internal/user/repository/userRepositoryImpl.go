package repository

import (
	"context"
	"errors"

	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/lib/pq"
)

type userPGRepository struct {
	db database.Database
}

func NewUserPGRepository(db database.Database) UserRepository {
	return &userPGRepository{db: db}
}

func (r *userPGRepository) CreateUser(ctx context.Context, user *dto.CreateUserDTO) (int, error) {
	row, err := r.db.GetDb().NamedQueryContext(ctx, createUserQuery, user)

	if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
		return -1, pgErr
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		return -1, err
	}

	var id int

	if row.Next() {
		row.Scan(&id)
	}

	return id, nil
}
