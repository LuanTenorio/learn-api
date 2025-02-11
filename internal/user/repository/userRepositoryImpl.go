package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/exception"
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

	//check unique constraint
	if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
		return -1, exception.New("There is already a user with this email", http.StatusConflict)
	} else if ok {
		return -1, exception.New("Error in operation with the database", http.StatusConflict, pgErr.Error(), "pg errorr")
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		return -1, exception.NewCanceledRequest(err.Error())
	}

	var id int

	if row.Next() {
		row.Scan(&id)
	}

	return id, nil
}

func (r *userPGRepository) FindUserAndPwdByEmail(ctx context.Context, email string) (*dto.UserWithPwdDTO, error) {
	row, err := r.db.GetDb().NamedQueryContext(ctx, selectUserWithEmailByPwdQuery, map[string]interface{}{
		"email": email,
	})

	if pgErr, ok := err.(*pq.Error); ok {
		return nil, exception.New("Error in operation with the database", http.StatusConflict, pgErr.Error(), "pg errorr")
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		return nil, exception.NewCanceledRequest(err.Error())
	}

	user := new(dto.UserWithPwdDTO)

	if row.Next() {
		row.StructScan(&user)
	}

	return user, nil
}
