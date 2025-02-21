package usecase

import (
	"context"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

func (u *subjectUsecaseImpl) Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.SubjectEntity, exception.Exception) {
	err := u.checkIfThereIsASubjectWithThisName(ctx, subjectDto.Name)

	if err != nil {
		return nil, err
	}

	subject, err := u.subjectRepo.Create(ctx, subjectDto)

	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (u *subjectUsecaseImpl) checkIfThereIsASubjectWithThisName(ctx context.Context, name string) exception.Exception {
	exist, err := u.subjectRepo.ExistSubjectByName(ctx, name)

	if err != nil {
		return err
	}

	if exist {
		return exception.New("There is already a story with that name", http.StatusConflict)
	}

	return nil
}
