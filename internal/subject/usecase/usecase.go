package usecase

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

type SubjectUsecase interface {
	Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.SubjectEntity, exception.Exception)
}
