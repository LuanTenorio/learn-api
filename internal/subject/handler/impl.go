package handler

import (
	"github.com/LuanTenorio/learn-api/internal/subject/usecase"
)

type subjectHandlerImpl struct {
	usecase usecase.SubjectUsecase
}

func New(uc usecase.SubjectUsecase) SubjectHandler {
	return &subjectHandlerImpl{usecase: uc}
}
