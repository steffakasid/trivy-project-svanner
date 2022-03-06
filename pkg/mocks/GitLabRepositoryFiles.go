// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "github.com/xanzy/go-gitlab"
)

// GitLabRepositoryFiles is an autogenerated mock type for the GitLabRepositoryFiles type
type GitLabRepositoryFiles struct {
	mock.Mock
}

type GitLabRepositoryFiles_Expecter struct {
	mock *mock.Mock
}

func (_m *GitLabRepositoryFiles) EXPECT() *GitLabRepositoryFiles_Expecter {
	return &GitLabRepositoryFiles_Expecter{mock: &_m.Mock}
}

// GetRawFile provides a mock function with given fields: pid, fileName, opt, options
func (_m *GitLabRepositoryFiles) GetRawFile(pid interface{}, fileName string, opt *gitlab.GetRawFileOptions, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, fileName, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(interface{}, string, *gitlab.GetRawFileOptions, ...gitlab.RequestOptionFunc) []byte); ok {
		r0 = rf(pid, fileName, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(interface{}, string, *gitlab.GetRawFileOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, fileName, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, string, *gitlab.GetRawFileOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, fileName, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GitLabRepositoryFiles_GetRawFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRawFile'
type GitLabRepositoryFiles_GetRawFile_Call struct {
	*mock.Call
}

// GetRawFile is a helper method to define mock.On call
//  - pid interface{}
//  - fileName string
//  - opt *gitlab.GetRawFileOptions
//  - options ...gitlab.RequestOptionFunc
func (_e *GitLabRepositoryFiles_Expecter) GetRawFile(pid interface{}, fileName interface{}, opt interface{}, options ...interface{}) *GitLabRepositoryFiles_GetRawFile_Call {
	return &GitLabRepositoryFiles_GetRawFile_Call{Call: _e.mock.On("GetRawFile",
		append([]interface{}{pid, fileName, opt}, options...)...)}
}

func (_c *GitLabRepositoryFiles_GetRawFile_Call) Run(run func(pid interface{}, fileName string, opt *gitlab.GetRawFileOptions, options ...gitlab.RequestOptionFunc)) *GitLabRepositoryFiles_GetRawFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(string), args[2].(*gitlab.GetRawFileOptions), variadicArgs...)
	})
	return _c
}

func (_c *GitLabRepositoryFiles_GetRawFile_Call) Return(_a0 []byte, _a1 *gitlab.Response, _a2 error) *GitLabRepositoryFiles_GetRawFile_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}
