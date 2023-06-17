// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	context "context"
	domains "robinhood/internal/core/domains"

	mock "github.com/stretchr/testify/mock"
)

// BlogRepository is an autogenerated mock type for the BlogRepository type
type BlogRepository struct {
	mock.Mock
}

type BlogRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *BlogRepository) EXPECT() *BlogRepository_Expecter {
	return &BlogRepository_Expecter{mock: &_m.Mock}
}

// Archive provides a mock function with given fields: _a0, _a1
func (_m *BlogRepository) Archive(_a0 context.Context, _a1 *domains.ArchiveBlogRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domains.ArchiveBlogRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlogRepository_Archive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Archive'
type BlogRepository_Archive_Call struct {
	*mock.Call
}

// Archive is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *domains.ArchiveBlogRequest
func (_e *BlogRepository_Expecter) Archive(_a0 interface{}, _a1 interface{}) *BlogRepository_Archive_Call {
	return &BlogRepository_Archive_Call{Call: _e.mock.On("Archive", _a0, _a1)}
}

func (_c *BlogRepository_Archive_Call) Run(run func(_a0 context.Context, _a1 *domains.ArchiveBlogRequest)) *BlogRepository_Archive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domains.ArchiveBlogRequest))
	})
	return _c
}

func (_c *BlogRepository_Archive_Call) Return(_a0 error) *BlogRepository_Archive_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlogRepository_Archive_Call) RunAndReturn(run func(context.Context, *domains.ArchiveBlogRequest) error) *BlogRepository_Archive_Call {
	_c.Call.Return(run)
	return _c
}

// Count provides a mock function with given fields: _a0, _a1
func (_m *BlogRepository) Count(_a0 context.Context, _a1 *domains.PaginationOptions) (int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domains.PaginationOptions) (int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domains.PaginationOptions) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domains.PaginationOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlogRepository_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type BlogRepository_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *domains.PaginationOptions
func (_e *BlogRepository_Expecter) Count(_a0 interface{}, _a1 interface{}) *BlogRepository_Count_Call {
	return &BlogRepository_Count_Call{Call: _e.mock.On("Count", _a0, _a1)}
}

func (_c *BlogRepository_Count_Call) Run(run func(_a0 context.Context, _a1 *domains.PaginationOptions)) *BlogRepository_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domains.PaginationOptions))
	})
	return _c
}

func (_c *BlogRepository_Count_Call) Return(_a0 int64, _a1 error) *BlogRepository_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlogRepository_Count_Call) RunAndReturn(run func(context.Context, *domains.PaginationOptions) (int64, error)) *BlogRepository_Count_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *BlogRepository) Create(_a0 context.Context, _a1 *domains.CreateBlogRequest) (*domains.Blog, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domains.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domains.CreateBlogRequest) (*domains.Blog, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domains.CreateBlogRequest) *domains.Blog); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domains.CreateBlogRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlogRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type BlogRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *domains.CreateBlogRequest
func (_e *BlogRepository_Expecter) Create(_a0 interface{}, _a1 interface{}) *BlogRepository_Create_Call {
	return &BlogRepository_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *BlogRepository_Create_Call) Run(run func(_a0 context.Context, _a1 *domains.CreateBlogRequest)) *BlogRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domains.CreateBlogRequest))
	})
	return _c
}

func (_c *BlogRepository_Create_Call) Return(_a0 *domains.Blog, _a1 error) *BlogRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlogRepository_Create_Call) RunAndReturn(run func(context.Context, *domains.CreateBlogRequest) (*domains.Blog, error)) *BlogRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTx provides a mock function with given fields: _a0, _a1, _a2
func (_m *BlogRepository) CreateTx(_a0 context.Context, _a1 *domains.CreateBlogRequest, _a2 domains.CreateBlogFn) (*domains.PopulatedBlog, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *domains.PopulatedBlog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domains.CreateBlogRequest, domains.CreateBlogFn) (*domains.PopulatedBlog, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domains.CreateBlogRequest, domains.CreateBlogFn) *domains.PopulatedBlog); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.PopulatedBlog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domains.CreateBlogRequest, domains.CreateBlogFn) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlogRepository_CreateTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTx'
type BlogRepository_CreateTx_Call struct {
	*mock.Call
}

// CreateTx is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *domains.CreateBlogRequest
//   - _a2 domains.CreateBlogFn
func (_e *BlogRepository_Expecter) CreateTx(_a0 interface{}, _a1 interface{}, _a2 interface{}) *BlogRepository_CreateTx_Call {
	return &BlogRepository_CreateTx_Call{Call: _e.mock.On("CreateTx", _a0, _a1, _a2)}
}

func (_c *BlogRepository_CreateTx_Call) Run(run func(_a0 context.Context, _a1 *domains.CreateBlogRequest, _a2 domains.CreateBlogFn)) *BlogRepository_CreateTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domains.CreateBlogRequest), args[2].(domains.CreateBlogFn))
	})
	return _c
}

func (_c *BlogRepository_CreateTx_Call) Return(_a0 *domains.PopulatedBlog, _a1 error) *BlogRepository_CreateTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlogRepository_CreateTx_Call) RunAndReturn(run func(context.Context, *domains.CreateBlogRequest, domains.CreateBlogFn) (*domains.PopulatedBlog, error)) *BlogRepository_CreateTx_Call {
	_c.Call.Return(run)
	return _c
}

// GetPopulatedBlogByID provides a mock function with given fields: _a0, _a1
func (_m *BlogRepository) GetPopulatedBlogByID(_a0 context.Context, _a1 string) (*domains.PopulatedBlog, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domains.PopulatedBlog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domains.PopulatedBlog, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domains.PopulatedBlog); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domains.PopulatedBlog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlogRepository_GetPopulatedBlogByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPopulatedBlogByID'
type BlogRepository_GetPopulatedBlogByID_Call struct {
	*mock.Call
}

// GetPopulatedBlogByID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *BlogRepository_Expecter) GetPopulatedBlogByID(_a0 interface{}, _a1 interface{}) *BlogRepository_GetPopulatedBlogByID_Call {
	return &BlogRepository_GetPopulatedBlogByID_Call{Call: _e.mock.On("GetPopulatedBlogByID", _a0, _a1)}
}

func (_c *BlogRepository_GetPopulatedBlogByID_Call) Run(run func(_a0 context.Context, _a1 string)) *BlogRepository_GetPopulatedBlogByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *BlogRepository_GetPopulatedBlogByID_Call) Return(_a0 *domains.PopulatedBlog, _a1 error) *BlogRepository_GetPopulatedBlogByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlogRepository_GetPopulatedBlogByID_Call) RunAndReturn(run func(context.Context, string) (*domains.PopulatedBlog, error)) *BlogRepository_GetPopulatedBlogByID_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *BlogRepository) List(_a0 context.Context, _a1 *domains.PaginationOptions) ([]domains.PopulatedBlog, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []domains.PopulatedBlog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domains.PaginationOptions) ([]domains.PopulatedBlog, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domains.PaginationOptions) []domains.PopulatedBlog); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domains.PopulatedBlog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domains.PaginationOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlogRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type BlogRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *domains.PaginationOptions
func (_e *BlogRepository_Expecter) List(_a0 interface{}, _a1 interface{}) *BlogRepository_List_Call {
	return &BlogRepository_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *BlogRepository_List_Call) Run(run func(_a0 context.Context, _a1 *domains.PaginationOptions)) *BlogRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domains.PaginationOptions))
	})
	return _c
}

func (_c *BlogRepository_List_Call) Return(_a0 []domains.PopulatedBlog, _a1 error) *BlogRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlogRepository_List_Call) RunAndReturn(run func(context.Context, *domains.PaginationOptions) ([]domains.PopulatedBlog, error)) *BlogRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: _a0, _a1
func (_m *BlogRepository) UpdateStatus(_a0 context.Context, _a1 *domains.UpdateBlogStatusRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domains.UpdateBlogStatusRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlogRepository_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type BlogRepository_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *domains.UpdateBlogStatusRequest
func (_e *BlogRepository_Expecter) UpdateStatus(_a0 interface{}, _a1 interface{}) *BlogRepository_UpdateStatus_Call {
	return &BlogRepository_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", _a0, _a1)}
}

func (_c *BlogRepository_UpdateStatus_Call) Run(run func(_a0 context.Context, _a1 *domains.UpdateBlogStatusRequest)) *BlogRepository_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domains.UpdateBlogStatusRequest))
	})
	return _c
}

func (_c *BlogRepository_UpdateStatus_Call) Return(_a0 error) *BlogRepository_UpdateStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlogRepository_UpdateStatus_Call) RunAndReturn(run func(context.Context, *domains.UpdateBlogStatusRequest) error) *BlogRepository_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewBlogRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlogRepository creates a new instance of BlogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlogRepository(t mockConstructorTestingTNewBlogRepository) *BlogRepository {
	mock := &BlogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}