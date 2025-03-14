// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/LuanTenorio/learn-api/internal/subject/dto"
	entity "github.com/LuanTenorio/learn-api/internal/subject/entity"

	exception "github.com/LuanTenorio/learn-api/internal/exception"

	mock "github.com/stretchr/testify/mock"

	pagination "github.com/LuanTenorio/learn-api/internal/pagination"
)

// SubjectRepository is an autogenerated mock type for the SubjectRepository type
type SubjectRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, subjectDto
func (_m *SubjectRepository) Create(ctx context.Context, subjectDto *dto.CreateSubjectDTO) (*entity.Subject, exception.Exception) {
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

// ExistSubjectByName provides a mock function with given fields: ctx, name, userId
func (_m *SubjectRepository) ExistSubjectByName(ctx context.Context, name string, userId int) (bool, exception.Exception) {
	ret := _m.Called(ctx, name, userId)

	if len(ret) == 0 {
		panic("no return value specified for ExistSubjectByName")
	}

	var r0 bool
	var r1 exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, string, int) (bool, exception.Exception)); ok {
		return rf(ctx, name, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) bool); ok {
		r0 = rf(ctx, name, userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) exception.Exception); ok {
		r1 = rf(ctx, name, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(exception.Exception)
		}
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ctx, _a1, userId
func (_m *SubjectRepository) FindMany(ctx context.Context, _a1 pagination.Pagination, userId int) ([]entity.Subject, int, exception.Exception) {
	ret := _m.Called(ctx, _a1, userId)

	if len(ret) == 0 {
		panic("no return value specified for FindMany")
	}

	var r0 []entity.Subject
	var r1 int
	var r2 exception.Exception
	if rf, ok := ret.Get(0).(func(context.Context, pagination.Pagination, int) ([]entity.Subject, int, exception.Exception)); ok {
		return rf(ctx, _a1, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pagination.Pagination, int) []entity.Subject); ok {
		r0 = rf(ctx, _a1, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Subject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pagination.Pagination, int) int); ok {
		r1 = rf(ctx, _a1, userId)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, pagination.Pagination, int) exception.Exception); ok {
		r2 = rf(ctx, _a1, userId)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(exception.Exception)
		}
	}

	return r0, r1, r2
}

// NewSubjectRepository creates a new instance of SubjectRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubjectRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubjectRepository {
	mock := &SubjectRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
