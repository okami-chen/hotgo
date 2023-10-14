// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/addons/game/dao/internal"
)

// internalDeviceDao is internal type for wrapping internal DAO implements.
type internalDeviceDao = *internal.DeviceDao

// deviceDao is the data access object for table ue_device.
// You can define custom methods on it to extend its functionality as you wish.
type deviceDao struct {
	internalDeviceDao
}

var (
	// Device is globally public accessible object for table ue_device operations.
	Device = deviceDao{
		internal.NewDeviceDao(),
	}
)

// Fill with you ideas below.
