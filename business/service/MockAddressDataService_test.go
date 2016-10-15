// Automatically generated by MockGen. DO NOT EDIT!
// Source: data/contract/AddressDataServiceContract.go

package service_test

import (
	gomock "github.com/golang/mock/gomock"
	. "github.com/microbusinesses/AddressService/data/contract"
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

func (_m *MockAddressDataService) Create(tenantID system.UUID, applicationID system.UUID, address Address) (system.UUID, error) {
	ret := _m.ctrl.Call(_m, "Create", tenantID, applicationID, address)
	ret0, _ := ret[0].(system.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAddressDataServiceRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2)
}

func (_m *MockAddressDataService) Update(tenantID system.UUID, applicationID system.UUID, addressID system.UUID, address Address) error {
	ret := _m.ctrl.Call(_m, "Update", tenantID, applicationID, addressID, address)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAddressDataServiceRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1, arg2, arg3)
}

func (_m *MockAddressDataService) Read(tenantID system.UUID, applicationID system.UUID, addressID system.UUID, keys []string) (Address, error) {
	ret := _m.ctrl.Call(_m, "Read", tenantID, applicationID, addressID, keys)
	ret0, _ := ret[0].(Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAddressDataServiceRecorder) Read(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1, arg2, arg3)
}

func (_m *MockAddressDataService) ReadAll(tenantID system.UUID, applicationID system.UUID, addressID system.UUID) (Address, error) {
	ret := _m.ctrl.Call(_m, "ReadAll", tenantID, applicationID, addressID)
	ret0, _ := ret[0].(Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAddressDataServiceRecorder) ReadAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadAll", arg0, arg1, arg2)
}

func (_m *MockAddressDataService) Delete(tenantID system.UUID, applicationID system.UUID, addressID system.UUID) error {
	ret := _m.ctrl.Call(_m, "Delete", tenantID, applicationID, addressID)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAddressDataServiceRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1, arg2)
}
