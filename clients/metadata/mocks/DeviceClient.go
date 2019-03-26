// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// DeviceClient is an autogenerated mock type for the DeviceClient type
type DeviceClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: dev, ctx
func (_m *DeviceClient) Add(dev *models.Device, ctx context.Context) (string, error) {
	ret := _m.Called(dev, ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.Device, context.Context) string); ok {
		r0 = rf(dev, ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Device, context.Context) error); ok {
		r1 = rf(dev, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckForDevice provides a mock function with given fields: token, ctx
func (_m *DeviceClient) CheckForDevice(token string, ctx context.Context) (models.Device, error) {
	ret := _m.Called(token, ctx)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) models.Device); ok {
		r0 = rf(token, ctx)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(token, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id, ctx
func (_m *DeviceClient) Delete(id string, ctx context.Context) error {
	ret := _m.Called(id, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, context.Context) error); ok {
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByName provides a mock function with given fields: name, ctx
func (_m *DeviceClient) DeleteByName(name string, ctx context.Context) error {
	ret := _m.Called(name, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, context.Context) error); ok {
		r0 = rf(name, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Device provides a mock function with given fields: id, ctx
func (_m *DeviceClient) Device(id string, ctx context.Context) (models.Device, error) {
	ret := _m.Called(id, ctx)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) models.Device); ok {
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceForName provides a mock function with given fields: name, ctx
func (_m *DeviceClient) DeviceForName(name string, ctx context.Context) (models.Device, error) {
	ret := _m.Called(name, ctx)

	var r0 models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) models.Device); ok {
		r0 = rf(name, ctx)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(name, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Devices provides a mock function with given fields: ctx
func (_m *DeviceClient) Devices(ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(context.Context) []models.Device); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesByLabel provides a mock function with given fields: label, ctx
func (_m *DeviceClient) DevicesByLabel(label string, ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(label, ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) []models.Device); ok {
		r0 = rf(label, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(label, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesForAddressable provides a mock function with given fields: addressableid, ctx
func (_m *DeviceClient) DevicesForAddressable(addressableid string, ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(addressableid, ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) []models.Device); ok {
		r0 = rf(addressableid, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(addressableid, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesForAddressableByName provides a mock function with given fields: addressableName, ctx
func (_m *DeviceClient) DevicesForAddressableByName(addressableName string, ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(addressableName, ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) []models.Device); ok {
		r0 = rf(addressableName, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(addressableName, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesForProfile provides a mock function with given fields: profileid, ctx
func (_m *DeviceClient) DevicesForProfile(profileid string, ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(profileid, ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) []models.Device); ok {
		r0 = rf(profileid, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(profileid, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesForProfileByName provides a mock function with given fields: profileName, ctx
func (_m *DeviceClient) DevicesForProfileByName(profileName string, ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(profileName, ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) []models.Device); ok {
		r0 = rf(profileName, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(profileName, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesForService provides a mock function with given fields: serviceid, ctx
func (_m *DeviceClient) DevicesForService(serviceid string, ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(serviceid, ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) []models.Device); ok {
		r0 = rf(serviceid, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(serviceid, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DevicesForServiceByName provides a mock function with given fields: serviceName, ctx
func (_m *DeviceClient) DevicesForServiceByName(serviceName string, ctx context.Context) ([]models.Device, error) {
	ret := _m.Called(serviceName, ctx)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(string, context.Context) []models.Device); ok {
		r0 = rf(serviceName, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(serviceName, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: dev, ctx
func (_m *DeviceClient) Update(dev models.Device, ctx context.Context) error {
	ret := _m.Called(dev, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Device, context.Context) error); ok {
		r0 = rf(dev, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdminState provides a mock function with given fields: id, adminState, ctx
func (_m *DeviceClient) UpdateAdminState(id string, adminState string, ctx context.Context) error {
	ret := _m.Called(id, adminState, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, context.Context) error); ok {
		r0 = rf(id, adminState, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdminStateByName provides a mock function with given fields: name, adminState, ctx
func (_m *DeviceClient) UpdateAdminStateByName(name string, adminState string, ctx context.Context) error {
	ret := _m.Called(name, adminState, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, context.Context) error); ok {
		r0 = rf(name, adminState, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLastConnected provides a mock function with given fields: id, time, ctx
func (_m *DeviceClient) UpdateLastConnected(id string, time int64, ctx context.Context) error {
	ret := _m.Called(id, time, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, context.Context) error); ok {
		r0 = rf(id, time, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLastConnectedByName provides a mock function with given fields: name, time, ctx
func (_m *DeviceClient) UpdateLastConnectedByName(name string, time int64, ctx context.Context) error {
	ret := _m.Called(name, time, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, context.Context) error); ok {
		r0 = rf(name, time, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLastReported provides a mock function with given fields: id, time, ctx
func (_m *DeviceClient) UpdateLastReported(id string, time int64, ctx context.Context) error {
	ret := _m.Called(id, time, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, context.Context) error); ok {
		r0 = rf(id, time, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLastReportedByName provides a mock function with given fields: name, time, ctx
func (_m *DeviceClient) UpdateLastReportedByName(name string, time int64, ctx context.Context) error {
	ret := _m.Called(name, time, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64, context.Context) error); ok {
		r0 = rf(name, time, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOpState provides a mock function with given fields: id, opState, ctx
func (_m *DeviceClient) UpdateOpState(id string, opState string, ctx context.Context) error {
	ret := _m.Called(id, opState, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, context.Context) error); ok {
		r0 = rf(id, opState, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOpStateByName provides a mock function with given fields: name, opState, ctx
func (_m *DeviceClient) UpdateOpStateByName(name string, opState string, ctx context.Context) error {
	ret := _m.Called(name, opState, ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, context.Context) error); ok {
		r0 = rf(name, opState, ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
