package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/internal/user/entity"
	"github.com/lib/pq"
)

type userPGRepository struct {
	db database.Database
}

func NewUserPGRepository(db database.Database) UserRepository {
	return &userPGRepository{db: db}
}

func (r *userPGRepository) CreateUser(ctx context.Context, userDto *dto.CreateUserDTO) (*entity.User, error) {
	row, err := r.db.GetDb().NamedQueryContext(ctx, createUserQuery, userDto)

	//check unique constraint
	if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
		return nil, exception.New("There is already a user with this email", http.StatusConflict)
	} else if ok {
		return nil, exception.New("Error in operation with the database", http.StatusConflict, pgErr.Error(), "pg errorr")
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		return nil, exception.NewCanceledRequest(err.Error())
	}

	user := entity.NewUserByCreateDto(userDto, "", 0)

	for row.Next() {
		row.StructScan(user)
	}

	return user, nil
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
		row.StructScan(user)
	} else {
		return nil, exception.New("There is no user with this email", http.StatusNotFound)
	}

	return user, nil
}
