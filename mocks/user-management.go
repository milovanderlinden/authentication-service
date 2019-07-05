// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/user-management-api.pb.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	api "github.com/influenzanet/authentication-service/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockUserManagementApiClient is a mock of UserManagementApiClient interface
type MockUserManagementApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserManagementApiClientMockRecorder
}

// MockUserManagementApiClientMockRecorder is the mock recorder for MockUserManagementApiClient
type MockUserManagementApiClientMockRecorder struct {
	mock *MockUserManagementApiClient
}

// NewMockUserManagementApiClient creates a new mock instance
func NewMockUserManagementApiClient(ctrl *gomock.Controller) *MockUserManagementApiClient {
	mock := &MockUserManagementApiClient{ctrl: ctrl}
	mock.recorder = &MockUserManagementApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserManagementApiClient) EXPECT() *MockUserManagementApiClientMockRecorder {
	return m.recorder
}

// Status mocks base method
func (m *MockUserManagementApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*api.Status, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockUserManagementApiClientMockRecorder) Status(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockUserManagementApiClient)(nil).Status), varargs...)
}

// LoginWithEmail mocks base method
func (m *MockUserManagementApiClient) LoginWithEmail(ctx context.Context, in *api.UserCredentials, opts ...grpc.CallOption) (*api.UserAuthInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoginWithEmail", varargs...)
	ret0, _ := ret[0].(*api.UserAuthInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginWithEmail indicates an expected call of LoginWithEmail
func (mr *MockUserManagementApiClientMockRecorder) LoginWithEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWithEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).LoginWithEmail), varargs...)
}

// SignupWithEmail mocks base method
func (m *MockUserManagementApiClient) SignupWithEmail(ctx context.Context, in *api.UserCredentials, opts ...grpc.CallOption) (*api.UserAuthInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignupWithEmail", varargs...)
	ret0, _ := ret[0].(*api.UserAuthInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupWithEmail indicates an expected call of SignupWithEmail
func (mr *MockUserManagementApiClientMockRecorder) SignupWithEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupWithEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).SignupWithEmail), varargs...)
}

// CheckRefreshToken mocks base method
func (m *MockUserManagementApiClient) CheckRefreshToken(ctx context.Context, in *api.UserReference, opts ...grpc.CallOption) (*api.Status, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckRefreshToken", varargs...)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRefreshToken indicates an expected call of CheckRefreshToken
func (mr *MockUserManagementApiClientMockRecorder) CheckRefreshToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRefreshToken", reflect.TypeOf((*MockUserManagementApiClient)(nil).CheckRefreshToken), varargs...)
}

// TokenRefreshed mocks base method
func (m *MockUserManagementApiClient) TokenRefreshed(ctx context.Context, in *api.UserReference, opts ...grpc.CallOption) (*api.Status, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TokenRefreshed", varargs...)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenRefreshed indicates an expected call of TokenRefreshed
func (mr *MockUserManagementApiClientMockRecorder) TokenRefreshed(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenRefreshed", reflect.TypeOf((*MockUserManagementApiClient)(nil).TokenRefreshed), varargs...)
}

// GetUser mocks base method
func (m *MockUserManagementApiClient) GetUser(ctx context.Context, in *api.UserReference, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserManagementApiClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserManagementApiClient)(nil).GetUser), varargs...)
}

// ChangePassword mocks base method
func (m *MockUserManagementApiClient) ChangePassword(ctx context.Context, in *api.PasswordChangeMsg, opts ...grpc.CallOption) (*api.Status, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangePassword", varargs...)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword
func (mr *MockUserManagementApiClientMockRecorder) ChangePassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserManagementApiClient)(nil).ChangePassword), varargs...)
}

// ChangeEmail mocks base method
func (m *MockUserManagementApiClient) ChangeEmail(ctx context.Context, in *api.EmailChangeMsg, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeEmail", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEmail indicates an expected call of ChangeEmail
func (mr *MockUserManagementApiClientMockRecorder) ChangeEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEmail", reflect.TypeOf((*MockUserManagementApiClient)(nil).ChangeEmail), varargs...)
}

// UpdateName mocks base method
func (m *MockUserManagementApiClient) UpdateName(ctx context.Context, in *api.NameUpdateRequest, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateName", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateName indicates an expected call of UpdateName
func (mr *MockUserManagementApiClientMockRecorder) UpdateName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockUserManagementApiClient)(nil).UpdateName), varargs...)
}

// DeleteAccount mocks base method
func (m *MockUserManagementApiClient) DeleteAccount(ctx context.Context, in *api.UserReference, opts ...grpc.CallOption) (*api.Status, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAccount", varargs...)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockUserManagementApiClientMockRecorder) DeleteAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockUserManagementApiClient)(nil).DeleteAccount), varargs...)
}

// UpdateBirthDate mocks base method
func (m *MockUserManagementApiClient) UpdateBirthDate(ctx context.Context, in *api.ProfileRequest, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBirthDate", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBirthDate indicates an expected call of UpdateBirthDate
func (mr *MockUserManagementApiClientMockRecorder) UpdateBirthDate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBirthDate", reflect.TypeOf((*MockUserManagementApiClient)(nil).UpdateBirthDate), varargs...)
}

// UpdateChildren mocks base method
func (m *MockUserManagementApiClient) UpdateChildren(ctx context.Context, in *api.ProfileRequest, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateChildren", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateChildren indicates an expected call of UpdateChildren
func (mr *MockUserManagementApiClientMockRecorder) UpdateChildren(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChildren", reflect.TypeOf((*MockUserManagementApiClient)(nil).UpdateChildren), varargs...)
}

// AddSubprofile mocks base method
func (m *MockUserManagementApiClient) AddSubprofile(ctx context.Context, in *api.SubProfileRequest, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddSubprofile", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSubprofile indicates an expected call of AddSubprofile
func (mr *MockUserManagementApiClientMockRecorder) AddSubprofile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubprofile", reflect.TypeOf((*MockUserManagementApiClient)(nil).AddSubprofile), varargs...)
}

// EditSubprofile mocks base method
func (m *MockUserManagementApiClient) EditSubprofile(ctx context.Context, in *api.SubProfileRequest, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditSubprofile", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditSubprofile indicates an expected call of EditSubprofile
func (mr *MockUserManagementApiClientMockRecorder) EditSubprofile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditSubprofile", reflect.TypeOf((*MockUserManagementApiClient)(nil).EditSubprofile), varargs...)
}

// RemoveSubprofile mocks base method
func (m *MockUserManagementApiClient) RemoveSubprofile(ctx context.Context, in *api.SubProfileRequest, opts ...grpc.CallOption) (*api.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveSubprofile", varargs...)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSubprofile indicates an expected call of RemoveSubprofile
func (mr *MockUserManagementApiClientMockRecorder) RemoveSubprofile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubprofile", reflect.TypeOf((*MockUserManagementApiClient)(nil).RemoveSubprofile), varargs...)
}

// MockUserManagementApiServer is a mock of UserManagementApiServer interface
type MockUserManagementApiServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserManagementApiServerMockRecorder
}

// MockUserManagementApiServerMockRecorder is the mock recorder for MockUserManagementApiServer
type MockUserManagementApiServerMockRecorder struct {
	mock *MockUserManagementApiServer
}

// NewMockUserManagementApiServer creates a new mock instance
func NewMockUserManagementApiServer(ctrl *gomock.Controller) *MockUserManagementApiServer {
	mock := &MockUserManagementApiServer{ctrl: ctrl}
	mock.recorder = &MockUserManagementApiServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserManagementApiServer) EXPECT() *MockUserManagementApiServerMockRecorder {
	return m.recorder
}

// Status mocks base method
func (m *MockUserManagementApiServer) Status(arg0 context.Context, arg1 *empty.Empty) (*api.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockUserManagementApiServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockUserManagementApiServer)(nil).Status), arg0, arg1)
}

// LoginWithEmail mocks base method
func (m *MockUserManagementApiServer) LoginWithEmail(arg0 context.Context, arg1 *api.UserCredentials) (*api.UserAuthInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginWithEmail", arg0, arg1)
	ret0, _ := ret[0].(*api.UserAuthInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginWithEmail indicates an expected call of LoginWithEmail
func (mr *MockUserManagementApiServerMockRecorder) LoginWithEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWithEmail", reflect.TypeOf((*MockUserManagementApiServer)(nil).LoginWithEmail), arg0, arg1)
}

// SignupWithEmail mocks base method
func (m *MockUserManagementApiServer) SignupWithEmail(arg0 context.Context, arg1 *api.UserCredentials) (*api.UserAuthInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignupWithEmail", arg0, arg1)
	ret0, _ := ret[0].(*api.UserAuthInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignupWithEmail indicates an expected call of SignupWithEmail
func (mr *MockUserManagementApiServerMockRecorder) SignupWithEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignupWithEmail", reflect.TypeOf((*MockUserManagementApiServer)(nil).SignupWithEmail), arg0, arg1)
}

// CheckRefreshToken mocks base method
func (m *MockUserManagementApiServer) CheckRefreshToken(arg0 context.Context, arg1 *api.UserReference) (*api.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRefreshToken indicates an expected call of CheckRefreshToken
func (mr *MockUserManagementApiServerMockRecorder) CheckRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRefreshToken", reflect.TypeOf((*MockUserManagementApiServer)(nil).CheckRefreshToken), arg0, arg1)
}

// TokenRefreshed mocks base method
func (m *MockUserManagementApiServer) TokenRefreshed(arg0 context.Context, arg1 *api.UserReference) (*api.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenRefreshed", arg0, arg1)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenRefreshed indicates an expected call of TokenRefreshed
func (mr *MockUserManagementApiServerMockRecorder) TokenRefreshed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenRefreshed", reflect.TypeOf((*MockUserManagementApiServer)(nil).TokenRefreshed), arg0, arg1)
}

// GetUser mocks base method
func (m *MockUserManagementApiServer) GetUser(arg0 context.Context, arg1 *api.UserReference) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserManagementApiServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserManagementApiServer)(nil).GetUser), arg0, arg1)
}

// ChangePassword mocks base method
func (m *MockUserManagementApiServer) ChangePassword(arg0 context.Context, arg1 *api.PasswordChangeMsg) (*api.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword
func (mr *MockUserManagementApiServerMockRecorder) ChangePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserManagementApiServer)(nil).ChangePassword), arg0, arg1)
}

// ChangeEmail mocks base method
func (m *MockUserManagementApiServer) ChangeEmail(arg0 context.Context, arg1 *api.EmailChangeMsg) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeEmail", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEmail indicates an expected call of ChangeEmail
func (mr *MockUserManagementApiServerMockRecorder) ChangeEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEmail", reflect.TypeOf((*MockUserManagementApiServer)(nil).ChangeEmail), arg0, arg1)
}

// UpdateName mocks base method
func (m *MockUserManagementApiServer) UpdateName(arg0 context.Context, arg1 *api.NameUpdateRequest) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateName", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateName indicates an expected call of UpdateName
func (mr *MockUserManagementApiServerMockRecorder) UpdateName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateName", reflect.TypeOf((*MockUserManagementApiServer)(nil).UpdateName), arg0, arg1)
}

// DeleteAccount mocks base method
func (m *MockUserManagementApiServer) DeleteAccount(arg0 context.Context, arg1 *api.UserReference) (*api.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(*api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockUserManagementApiServerMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockUserManagementApiServer)(nil).DeleteAccount), arg0, arg1)
}

// UpdateBirthDate mocks base method
func (m *MockUserManagementApiServer) UpdateBirthDate(arg0 context.Context, arg1 *api.ProfileRequest) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBirthDate", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBirthDate indicates an expected call of UpdateBirthDate
func (mr *MockUserManagementApiServerMockRecorder) UpdateBirthDate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBirthDate", reflect.TypeOf((*MockUserManagementApiServer)(nil).UpdateBirthDate), arg0, arg1)
}

// UpdateChildren mocks base method
func (m *MockUserManagementApiServer) UpdateChildren(arg0 context.Context, arg1 *api.ProfileRequest) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChildren", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateChildren indicates an expected call of UpdateChildren
func (mr *MockUserManagementApiServerMockRecorder) UpdateChildren(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChildren", reflect.TypeOf((*MockUserManagementApiServer)(nil).UpdateChildren), arg0, arg1)
}

// AddSubprofile mocks base method
func (m *MockUserManagementApiServer) AddSubprofile(arg0 context.Context, arg1 *api.SubProfileRequest) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubprofile", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSubprofile indicates an expected call of AddSubprofile
func (mr *MockUserManagementApiServerMockRecorder) AddSubprofile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubprofile", reflect.TypeOf((*MockUserManagementApiServer)(nil).AddSubprofile), arg0, arg1)
}

// EditSubprofile mocks base method
func (m *MockUserManagementApiServer) EditSubprofile(arg0 context.Context, arg1 *api.SubProfileRequest) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditSubprofile", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditSubprofile indicates an expected call of EditSubprofile
func (mr *MockUserManagementApiServerMockRecorder) EditSubprofile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditSubprofile", reflect.TypeOf((*MockUserManagementApiServer)(nil).EditSubprofile), arg0, arg1)
}

// RemoveSubprofile mocks base method
func (m *MockUserManagementApiServer) RemoveSubprofile(arg0 context.Context, arg1 *api.SubProfileRequest) (*api.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSubprofile", arg0, arg1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSubprofile indicates an expected call of RemoveSubprofile
func (mr *MockUserManagementApiServerMockRecorder) RemoveSubprofile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubprofile", reflect.TypeOf((*MockUserManagementApiServer)(nil).RemoveSubprofile), arg0, arg1)
}
