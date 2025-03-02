package pagination

import "github.com/LuanTenorio/learn-api/internal/pagination/dto"

type Pagination interface {
	Offset() int
	Limit() int
	NewResponse(data interface{}, tot int) *dto.PaginationResponseDTO
	Validate() bool
}
