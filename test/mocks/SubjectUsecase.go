// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/LuanTenorio/learn-api/internal/subject/dto"
	entity "github.com/LuanTenorio/learn-api/internal/subject/entity"

	exception "github.com/LuanTenorio/learn-api/internal/exception"

	mock "github.com/stretchr/testify/mock"

	pagination "github.com/LuanTenorio/learn-api/internal/pagination"

	paginationdto "github.com/LuanTenorio/learn-api/internal/pagination/dto"
)

// SubjectUsecase is an autogenerated mock type for the SubjectUsecase type
type SubjectUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, subjectDto
func (_m *SubjectUsecase) Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception) {
	ret := _m.Called(ctx, subjectDto)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *entity.Subject
	var r1 exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception)); ok {
		return rf(ctx, subjectDto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateSubjectDTO) *entity.Subject); ok {
		r0 = rf(ctx, subjectDto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Subject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreateSubjectDTO) exception.Exception); ok {
		r1 = rf(ctx, subjectDto)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(exception.Exception)
		}
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ctx, _a1, userId
func (_m *SubjectUsecase) FindMany(ctx context.Context, _a1 pagination.Pagination, userId int) (*paginationdto.PaginationResponseDTO, exception.Exception) {
	ret := _m.Called(ctx, _a1, userId)

	if len(ret) == 0 {
		panic("no return value specified for FindMany")
	}

	var r0 *paginationdto.PaginationResponseDTO
	var r1 exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, pagination.Pagination, int) (*paginationdto.PaginationResponseDTO, exception.Exception)); ok {
		return rf(ctx, _a1, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pagination.Pagination, int) *paginationdto.PaginationResponseDTO); ok {
		r0 = rf(ctx, _a1, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paginationdto.PaginationResponseDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pagination.Pagination, int) exception.Exception); ok {
		r1 = rf(ctx, _a1, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(exception.Exception)
		}
	}

	return r0, r1
}

// NewSubjectUsecase creates a new instance of SubjectUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubjectUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubjectUsecase {
	mock := &SubjectUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
