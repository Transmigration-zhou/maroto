// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	metrics "github.com/johnfercher/maroto/v2/pkg/metrics"
	mock "github.com/stretchr/testify/mock"
)

// Document is an autogenerated mock type for the Document type
type Document struct {
	mock.Mock
}

type Document_Expecter struct {
	mock *mock.Mock
}

func (_m *Document) EXPECT() *Document_Expecter {
	return &Document_Expecter{mock: &_m.Mock}
}

// GetBase64 provides a mock function with given fields:
func (_m *Document) GetBase64() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Document_GetBase64_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBase64'
type Document_GetBase64_Call struct {
	*mock.Call
}

// GetBase64 is a helper method to define mock.On call
func (_e *Document_Expecter) GetBase64() *Document_GetBase64_Call {
	return &Document_GetBase64_Call{Call: _e.mock.On("GetBase64")}
}

func (_c *Document_GetBase64_Call) Run(run func()) *Document_GetBase64_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Document_GetBase64_Call) Return(_a0 string) *Document_GetBase64_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Document_GetBase64_Call) RunAndReturn(run func() string) *Document_GetBase64_Call {
	_c.Call.Return(run)
	return _c
}

// GetBytes provides a mock function with given fields:
func (_m *Document) GetBytes() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Document_GetBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBytes'
type Document_GetBytes_Call struct {
	*mock.Call
}

// GetBytes is a helper method to define mock.On call
func (_e *Document_Expecter) GetBytes() *Document_GetBytes_Call {
	return &Document_GetBytes_Call{Call: _e.mock.On("GetBytes")}
}

func (_c *Document_GetBytes_Call) Run(run func()) *Document_GetBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Document_GetBytes_Call) Return(_a0 []byte) *Document_GetBytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Document_GetBytes_Call) RunAndReturn(run func() []byte) *Document_GetBytes_Call {
	_c.Call.Return(run)
	return _c
}

// GetReport provides a mock function with given fields:
func (_m *Document) GetReport() *metrics.Report {
	ret := _m.Called()

	var r0 *metrics.Report
	if rf, ok := ret.Get(0).(func() *metrics.Report); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metrics.Report)
		}
	}

	return r0
}

// Document_GetReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReport'
type Document_GetReport_Call struct {
	*mock.Call
}

// GetReport is a helper method to define mock.On call
func (_e *Document_Expecter) GetReport() *Document_GetReport_Call {
	return &Document_GetReport_Call{Call: _e.mock.On("GetReport")}
}

func (_c *Document_GetReport_Call) Run(run func()) *Document_GetReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Document_GetReport_Call) Return(_a0 *metrics.Report) *Document_GetReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Document_GetReport_Call) RunAndReturn(run func() *metrics.Report) *Document_GetReport_Call {
	_c.Call.Return(run)
	return _c
}

// Merge provides a mock function with given fields: _a0
func (_m *Document) Merge(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Document_Merge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Merge'
type Document_Merge_Call struct {
	*mock.Call
}

// Merge is a helper method to define mock.On call
//   - _a0 []byte
func (_e *Document_Expecter) Merge(_a0 interface{}) *Document_Merge_Call {
	return &Document_Merge_Call{Call: _e.mock.On("Merge", _a0)}
}

func (_c *Document_Merge_Call) Run(run func(_a0 []byte)) *Document_Merge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *Document_Merge_Call) Return(_a0 error) *Document_Merge_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Document_Merge_Call) RunAndReturn(run func([]byte) error) *Document_Merge_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: file
func (_m *Document) Save(file string) error {
	ret := _m.Called(file)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Document_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Document_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - file string
func (_e *Document_Expecter) Save(file interface{}) *Document_Save_Call {
	return &Document_Save_Call{Call: _e.mock.On("Save", file)}
}

func (_c *Document_Save_Call) Run(run func(file string)) *Document_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Document_Save_Call) Return(_a0 error) *Document_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Document_Save_Call) RunAndReturn(run func(string) error) *Document_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewDocument creates a new instance of Document. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDocument(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Document {
	mock := &Document{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
