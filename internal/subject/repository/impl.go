package repository

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/pagination"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
	"github.com/jmoiron/sqlx"
)

type SubjectPgRepository struct {
	db database.Database
}

func New(db database.Database) SubjectRepository {
	return &SubjectPgRepository{db: db}
}

func (r *SubjectPgRepository) Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception) {
	row, err := r.db.GetDb().NamedQueryContext(ctx, createSubjectQuery, subjectDto)

	subject := new(entity.Subject)

	if pgErr := exception.CheckDbException(err); pgErr != nil {
		return nil, pgErr.AddTraceLog("Exception in the database")
	} else if scanErr := database.StructScanOrError(row, subject); scanErr != nil {
		return nil, scanErr.AddTraceLog("Erro on parse subjcet struct")
	}

	return subject, nil
}

func (r *SubjectPgRepository) ExistSubjectByName(ctx context.Context, name string, userId int) (bool, exception.Exception) {
	var id int
	err := r.db.GetDb().GetContext(ctx, &id, findIdByNameQuery, name, userId)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if pgErr := exception.CheckDbException(err); pgErr != nil {
		return false, pgErr.AddTraceLog("Exception in the database")
	}

	return true, nil
}

func (r *SubjectPgRepository) FindMany(ctx context.Context, pagination pagination.Pagination, userId int) ([]entity.Subject, int, exception.Exception) {
	var subjects []entity.Subject

	tx, err := r.db.GetDb().Beginx()

	if err != nil {
		return nil, 0, exception.New("Internal db error", http.StatusInternalServerError, "Error when opening transaction", err.Error())
	}

	tot, totErr := findTotalItems(ctx, tx, userId)

	if totErr != nil {
		return nil, 0, totErr.AddTraceLog("Repository: Exception in the database")
	}

	err = tx.SelectContext(ctx, &subjects, findPaginationQuery, pagination.Limit(), pagination.Offset())

	if err == sql.ErrNoRows {
		return nil, 0, nil
	} else if pgErr := exception.CheckDbException(err); pgErr != nil {
		return nil, 0, pgErr.AddTraceLog("Repository: Exception in the database")
	}

	err = tx.Commit()

	if err != nil {
		return nil, 0, exception.New("Internal db error", http.StatusInternalServerError, "Error when commit transaction", err.Error())
	}

	return subjects, tot, nil
}

func findTotalItems(ctx context.Context, tx *sqlx.Tx, userId int) (int, exception.Exception) {
	var total int
	err := tx.GetContext(ctx, &total, totalItemsQuery, userId)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	if pgErr := exception.CheckDbException(err); pgErr != nil {
		return 0, pgErr.AddTraceLog("Exception in the database")
	}

	return total, nil
}
