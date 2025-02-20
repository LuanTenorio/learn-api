package repository

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

type SubjectRepository interface {
	Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.SubjectEntity, error)
	ExistSubjectByName(ctx context.Context, name string) (bool, error)
}
