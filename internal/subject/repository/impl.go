package repository

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

type SubjectPgRepository struct {
	db database.Database
}

func New(db database.Database) SubjectRepository {
	return &SubjectPgRepository{db: db}
}

func (r *SubjectPgRepository) Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.SubjectEntity, error) {
	row, err := r.db.GetDb().NamedQueryContext(ctx, createSubjectQuery, subjectDto)
	subject := new(entity.SubjectEntity)

	if pgErr := exception.CheckDbException(err); pgErr != nil {
		return nil, pgErr
	} else if scanErr := database.StructScanOrError(row, subject); scanErr != nil {
		return nil, scanErr
	}

	return subject, nil
}

func (r *SubjectPgRepository) ExistSubjectByName(ctx context.Context, name string) (bool, error) {
	var id int

	err := r.db.GetDb().GetContext(ctx, id, findIdByName, name)

	if pgErr := exception.CheckDbException(err); pgErr != nil {
		return false, pgErr
	}

	return id != 0, nil
}
