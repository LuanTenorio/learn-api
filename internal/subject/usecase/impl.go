package usecase

import "github.com/LuanTenorio/learn-api/internal/subject/repository"

type subjectUsecaseImpl struct {
	subjectRepo repository.SubjectRepository
}

func New(subjectRepo repository.SubjectRepository) SubjectUsecase {
	return &subjectUsecaseImpl{subjectRepo: subjectRepo}
}
