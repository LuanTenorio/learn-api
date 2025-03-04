package repository

import (
	"context"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/database/sqlc"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/internal/user/entity"
	"github.com/jackc/pgx/v5/pgconn"
)

type userPGRepository struct {
	db database.Database
}

func NewUserPGRepository(db database.Database) UserRepository {
	return &userPGRepository{db: db}
}

func (r *userPGRepository) CreateUser(ctx context.Context, u *dto.CreateUserDTO) (*entity.User, exception.Exception) {
	user, err := r.db.GetQueries().CreateUser(ctx, sqlc.CreateUserParams{Name: u.Name, Email: u.Email, Password: u.Password})

	//check unique constraint
	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
		return nil, exception.New("There is already a user with this email", http.StatusConflict)
	} else if pgErr := exception.CheckDbException(err); pgErr != nil {
		return nil, pgErr.AddTraceLog("Exception in the database")
	}

	return entity.M2E(&user), nil
}

func (r *userPGRepository) FindUserAndPwdByEmail(ctx context.Context, email string) (*dto.UserWithPwdDTO, exception.Exception) {
	user, err := r.db.GetQueries().FindUserByEmail(ctx, email)

	if err == database.ErrNoRows {
		return nil, exception.New("User with this email was not found", http.StatusNotFound, err.Error())
	} else if pgErr := exception.CheckDbException(err); pgErr != nil {
		return nil, pgErr.AddTraceLog("Exception in the database")
	}

	return &dto.UserWithPwdDTO{Id: int(user.ID), Name: user.Name, Email: user.Email, Password: user.Password, CreatedAt: user.CreatedAt.Time.String()}, nil
}
