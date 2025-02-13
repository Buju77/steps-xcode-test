// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	command "github.com/bitrise-io/go-utils/command"
	mock "github.com/stretchr/testify/mock"

	version "github.com/hashicorp/go-version"
)

// Xcpretty is an autogenerated mock type for the Xcpretty type
type Xcpretty struct {
	mock.Mock
}

// Install provides a mock function with given fields:
func (_m *Xcpretty) Install() ([]command.Command, error) {
	ret := _m.Called()

	var r0 []command.Command
	if rf, ok := ret.Get(0).(func() []command.Command); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]command.Command)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsInstalled provides a mock function with given fields:
func (_m *Xcpretty) IsInstalled() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields:
func (_m *Xcpretty) Version() (*version.Version, error) {
	ret := _m.Called()

	var r0 *version.Version
	if rf, ok := ret.Get(0).(func() *version.Version); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
