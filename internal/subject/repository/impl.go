package repository

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/database/sqlc"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/pagination"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

type SubjectPgRepository struct {
	db database.Database
}

func New(db database.Database) SubjectRepository {
	return &SubjectPgRepository{db: db}
}

func (r *SubjectPgRepository) Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception) {
	subject, err := r.db.GetQueries().CreateSubject(ctx, sqlc.CreateSubjectParams{Name: subjectDto.Name, UserID: int32(subjectDto.UserId)})

	if pgErr := exception.CheckDbException(err); pgErr != nil {
		return nil, pgErr.AddTraceLog("Exception in the database")
	}

	return entity.M2E(&subject), nil
}

func (r *SubjectPgRepository) ExistSubjectByName(ctx context.Context, name string, userId int) (bool, exception.Exception) {
	_, err := r.db.GetQueries().FindSubjectByIdAndName(ctx, sqlc.FindSubjectByIdAndNameParams{Name: name, UserID: 83})

	if err == database.ErrNoRows {
		return false, nil
	} else if pgErr := exception.CheckDbException(err); pgErr != nil {
		return false, pgErr.AddTraceLog("Exception in the database")
	}

	return true, nil
}

func (r *SubjectPgRepository) List(ctx context.Context, pagination pagination.Pagination, userId int) ([]*entity.Subject, int, exception.Exception) {
	tx, err := r.db.GetDb().Begin(ctx)

	if err != nil {
		return nil, 0, exception.NewDB(err.Error())
	}

	qtx := r.db.GetQueries().WithTx(tx)

	subjectsResp, listErr := findMany(ctx, qtx, pagination, userId)

	if listErr != nil {
		tx.Rollback(ctx)
		return nil, 0, listErr
	}

	tot, totErr := findTotalItems(ctx, qtx, userId)

	if totErr != nil {
		tx.Rollback(ctx)
		return nil, 0, totErr
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, 0, exception.NewDB(err.Error())
	}

	return subjectsResp, tot, nil
}

func findMany(ctx context.Context, qtx *sqlc.Queries, pagination pagination.Pagination, userId int) ([]*entity.Subject, exception.Exception) {
	paginationDto := sqlc.ListSubjectsParams{UserID: int32(userId), Limit: int32(pagination.Limit()), Offset: int32(pagination.Offset())}
	subjects, err := qtx.ListSubjects(ctx, paginationDto)

	if err == database.ErrNoRows {
		return nil, nil
	} else if pgErr := exception.CheckDbException(err); pgErr != nil {
		return nil, pgErr.AddTraceLog("Repository: Exception in the database")
	}

	subjectsResponse := make([]*entity.Subject, len(subjects))

	for i, s := range subjects {
		subjectsResponse[i] = entity.M2E(&s)
	}

	return subjectsResponse, nil
}

func findTotalItems(ctx context.Context, qtx *sqlc.Queries, userId int) (int, exception.Exception) {
	tot, err := qtx.TotalSubjectsByUser(ctx, int32(userId))

	if err != nil {
		return 0, exception.NewDB(err.Error())
	}

	return int(tot), nil
}
