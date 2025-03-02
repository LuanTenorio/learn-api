package usecase

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/pagination"
	"github.com/LuanTenorio/learn-api/internal/pagination/dto"
)

func (u *subjectUsecaseImpl) FindMany(ctx context.Context, pagination pagination.Pagination, userId int) (*dto.PaginationResponseDTO, exception.Exception) {
	subjects, tot, err := u.subjectRepo.FindMany(ctx, pagination, userId)

	if err != nil {
		return nil, err
	}

	return pagination.NewResponse(subjects, tot), nil
}
