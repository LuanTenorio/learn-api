package pagination

import (
	"math"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/auth"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/pagination/dto"
	"github.com/LuanTenorio/learn-api/internal/util"
	"github.com/labstack/echo/v4"
)

const (
	defaultLimit = 10
	defaultPage  = 1
)

type paginationImpl struct {
	page  int
	pages int
	limit int
}

func NewDefault() Pagination {
	return &paginationImpl{page: defaultPage, pages: 0, limit: defaultLimit}
}

func newByDTO(dto *dto.PaginationRequestDTO) Pagination {
	if dto.Limit == 0 {
		dto.Limit = defaultLimit
	}

	if dto.Page == 0 {
		dto.Page = defaultPage
	}

	return &paginationImpl{page: dto.Page, limit: dto.Limit, pages: 0}
}

func (p *paginationImpl) NewResponse(data interface{}, tot int) *dto.PaginationResponseDTO {
	totalPages := int(math.Ceil(float64(tot) / float64(p.limit)))
	return &dto.PaginationResponseDTO{Page: p.page, Limit: p.limit, Pages: totalPages, TotalItems: tot, Data: data}
}

func (p *paginationImpl) Validate() bool {
	return limitIsValid(p.limit) && p.page > 0
}

func limitIsValid(l int) bool {
	return l <= 50 && l > 0
}

func (p *paginationImpl) Offset() int {
	return (p.page - 1) * p.limit
}

func (p *paginationImpl) Limit() int {
	return p.limit
}

func GetPaginationFromParams(c echo.Context) (Pagination, exception.Exception) {
	paginationDto := new(dto.PaginationRequestDTO)

	if err := util.BindDataRequest(c, paginationDto); err != nil {
		return nil, err
	}

	pagination := newByDTO(paginationDto)

	if valid := pagination.Validate(); !valid {
		return nil, exception.New("Invalid pagination", http.StatusBadRequest)
	}

	return pagination, nil
}

func getUserId(c echo.Context) (int, exception.Exception) {
	if claims, ok := c.Get("claims").(*auth.JwtCustomClaims); ok {
		return claims.User.Id, nil
	}

	return 0, exception.New("Internal Error", http.StatusInternalServerError, "Error when taking the jwt id")
}

func GetPaginationAndUserId(c echo.Context) (Pagination, int, exception.Exception) {
	userId, err := getUserId(c)
	pgt, pgtErr := GetPaginationFromParams(c)

	if err == nil {
		err = pgtErr
	}

	return pgt, userId, err
}
