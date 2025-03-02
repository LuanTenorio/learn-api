package usecase

import (
	"context"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

func (u *subjectUsecaseImpl) Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception) {
	err := u.checkIfThereIsASubjectWithThisName(ctx, subjectDto)

	if err != nil {
		return nil, err
	}

	subject, err := u.subjectRepo.Create(ctx, subjectDto)

	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (u *subjectUsecaseImpl) checkIfThereIsASubjectWithThisName(ctx context.Context, subjectDto *dto.CreateSubjectDTO) exception.Exception {
	exist, err := u.subjectRepo.ExistSubjectByName(ctx, subjectDto.Name, subjectDto.UserId)

	if err != nil {
		return err
	}

	if exist {
		return exception.New("There is already a subject with that name", http.StatusConflict)
	}

	return nil
}
