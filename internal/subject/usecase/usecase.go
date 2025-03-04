package usecase

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/pagination"
	paginationDto "github.com/LuanTenorio/learn-api/internal/pagination/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/dto"
	"github.com/LuanTenorio/learn-api/internal/subject/entity"
)

type SubjectUsecase interface {
	Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception)
	List(ctx context.Context, pagination pagination.Pagination, userId int) (*paginationDto.PaginationResponseDTO, exception.Exception)
}
