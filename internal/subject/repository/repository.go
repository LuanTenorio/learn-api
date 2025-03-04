package repository

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/pagination"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

type SubjectRepository interface {
	Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception)
	ExistSubjectByName(ctx context.Context, name string, userId int) (bool, exception.Exception)
	List(ctx context.Context, pagination pagination.Pagination, userId int) ([]*entity.Subject, int, exception.Exception)
}
