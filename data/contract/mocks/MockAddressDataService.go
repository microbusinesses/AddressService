// Automatically generated by MockGen. DO NOT EDIT!
// Source: AddressDataServiceContract.go

package mock_contract

import (
	gomock "github.com/golang/mock/gomock"
	shared "github.com/microbusinesses/AddressService/data/shared"
	system "github.com/microbusinesses/Micro-Businesses-Core/system"
)

// Mock of AddressDataService interface
type MockAddressDataService struct {
	ctrl     *gomock.Controller
	recorder *_MockAddressDataServiceRecorder
}

// Recorder for MockAddressDataService (not exported)
type _MockAddressDataServiceRecorder struct {
	mock *MockAddressDataService
}

func NewMockAddressDataService(ctrl *gomock.Controller) *MockAddressDataService {
	mock := &MockAddressDataService{ctrl: ctrl}
	mock.recorder = &_MockAddressDataServiceRecorder{mock}
	return mock
}

func (_m *MockAddressDataService) EXPECT() *_MockAddressDataServiceRecorder {
	return _m.recorder
}

func (_m *MockAddressDataService) Create(tenantId system.UUID, applicationId system.UUID, address shared.Address) (system.UUID, error) {
	ret := _m.ctrl.Call(_m, "Create", tenantId, applicationId, address)
	ret0, _ := ret[0].(system.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAddressDataServiceRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2)
}

func (_m *MockAddressDataService) Update(tenantId system.UUID, applicationId system.UUID, addressId system.UUID, address shared.Address) error {
	ret := _m.ctrl.Call(_m, "Update", tenantId, applicationId, addressId, address)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAddressDataServiceRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1, arg2, arg3)
}

func (_m *MockAddressDataService) Read(tenantId system.UUID, applicationId system.UUID, addressId system.UUID, detailsKeys []string) (shared.Address, error) {
	ret := _m.ctrl.Call(_m, "Read", tenantId, applicationId, addressId, detailsKeys)
	ret0, _ := ret[0].(shared.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAddressDataServiceRecorder) Read(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1, arg2, arg3)
}

func (_m *MockAddressDataService) ReadAll(tenantId system.UUID, applicationId system.UUID, addressId system.UUID) (shared.Address, error) {
	ret := _m.ctrl.Call(_m, "ReadAll", tenantId, applicationId, addressId)
	ret0, _ := ret[0].(shared.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAddressDataServiceRecorder) ReadAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadAll", arg0, arg1, arg2)
}

func (_m *MockAddressDataService) Delete(tenantId system.UUID, applicationId system.UUID, addressId system.UUID) error {
	ret := _m.ctrl.Call(_m, "Delete", tenantId, applicationId, addressId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAddressDataServiceRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1, arg2)
}
