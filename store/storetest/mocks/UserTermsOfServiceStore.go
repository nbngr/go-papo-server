// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "bitbucket.org/enesyteam/papo-server/model"
	mock "github.com/stretchr/testify/mock"
)

// UserTermsOfServiceStore is an autogenerated mock type for the UserTermsOfServiceStore type
type UserTermsOfServiceStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userId, termsOfServiceId
func (_m *UserTermsOfServiceStore) Delete(userId string, termsOfServiceId string) error {
	ret := _m.Called(userId, termsOfServiceId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, termsOfServiceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUser provides a mock function with given fields: userId
func (_m *UserTermsOfServiceStore) GetByUser(userId string) (*model.UserTermsOfService, error) {
	ret := _m.Called(userId)

	var r0 *model.UserTermsOfService
	if rf, ok := ret.Get(0).(func(string) *model.UserTermsOfService); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserTermsOfService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: userTermsOfService
func (_m *UserTermsOfServiceStore) Save(userTermsOfService *model.UserTermsOfService) (*model.UserTermsOfService, error) {
	ret := _m.Called(userTermsOfService)

	var r0 *model.UserTermsOfService
	if rf, ok := ret.Get(0).(func(*model.UserTermsOfService) *model.UserTermsOfService); ok {
		r0 = rf(userTermsOfService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserTermsOfService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserTermsOfService) error); ok {
		r1 = rf(userTermsOfService)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
