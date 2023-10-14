// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAdminNoticeDao is internal type for wrapping internal DAO implements.
type internalAdminNoticeDao = *internal.AdminNoticeDao

// adminNoticeDao is the data access object for table hg_admin_notice.
// You can define custom methods on it to extend its functionality as you wish.
type adminNoticeDao struct {
	internalAdminNoticeDao
}

var (
	// AdminNotice is globally common accessible object for table hg_admin_notice operations.
	AdminNotice = adminNoticeDao{
		internal.NewAdminNoticeDao(),
	}
)

// Fill with you ideas below.
