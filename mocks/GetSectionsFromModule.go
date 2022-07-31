// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/jeanmolossi/effective-eureka/src/core/sections/domain"
	mock "github.com/stretchr/testify/mock"
)

// GetSectionsFromModule is an autogenerated mock type for the GetSectionsFromModule type
type GetSectionsFromModule struct {
	mock.Mock
}

// Run provides a mock function with given fields: input
func (_m *GetSectionsFromModule) Run(input *domain.GetSectionsParams) ([]domain.Section, error) {
	ret := _m.Called(input)

	var r0 []domain.Section
	if rf, ok := ret.Get(0).(func(*domain.GetSectionsParams) []domain.Section); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.GetSectionsParams) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
