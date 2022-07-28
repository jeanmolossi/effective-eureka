// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Module is an autogenerated mock type for the Module type
type Module struct {
	mock.Mock
}

// GetCourseID provides a mock function with given fields:
func (_m *Module) GetCourseID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetModuleDescription provides a mock function with given fields:
func (_m *Module) GetModuleDescription() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetModuleID provides a mock function with given fields:
func (_m *Module) GetModuleID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetModuleThumb provides a mock function with given fields:
func (_m *Module) GetModuleThumb() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetModuleTitle provides a mock function with given fields:
func (_m *Module) GetModuleTitle() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsModulePublished provides a mock function with given fields:
func (_m *Module) IsModulePublished() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PublishModule provides a mock function with given fields:
func (_m *Module) PublishModule() {
	_m.Called()
}

// SetCourseID provides a mock function with given fields: courseID
func (_m *Module) SetCourseID(courseID string) {
	_m.Called(courseID)
}

// SetModuleDescription provides a mock function with given fields: desc
func (_m *Module) SetModuleDescription(desc string) {
	_m.Called(desc)
}

// SetModuleID provides a mock function with given fields: moduleID
func (_m *Module) SetModuleID(moduleID string) {
	_m.Called(moduleID)
}

// SetModuleThumb provides a mock function with given fields: thumb
func (_m *Module) SetModuleThumb(thumb string) {
	_m.Called(thumb)
}

// SetModuleTitle provides a mock function with given fields: title
func (_m *Module) SetModuleTitle(title string) {
	_m.Called(title)
}

// UnpublishModule provides a mock function with given fields:
func (_m *Module) UnpublishModule() {
	_m.Called()
}