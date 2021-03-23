// Code generated by mockery v2.3.0. DO NOT EDIT.

package storagemocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	user "github.com/vosgaust/voicemod-challenge.git/internal/domain/user"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, userID
func (_m *UserRepository) Delete(ctx context.Context, userID user.UserID) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, user.UserID) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, userID
func (_m *UserRepository) Find(ctx context.Context, userID user.UserID) (*user.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(context.Context, user.UserID) *user.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, user.UserID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, userEmail
func (_m *UserRepository) FindByEmail(ctx context.Context, userEmail user.UserEmail) (*user.User, error) {
	ret := _m.Called(ctx, userEmail)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(context.Context, user.UserEmail) *user.User); ok {
		r0 = rf(ctx, userEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, user.UserEmail) error); ok {
		r1 = rf(ctx, userEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) Save(ctx context.Context, _a1 user.User) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, user.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *UserRepository) Update(ctx context.Context, _a1 user.User) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, user.User) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
