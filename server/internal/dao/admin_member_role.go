// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAdminMemberRoleDao is internal type for wrapping internal DAO implements.
type internalAdminMemberRoleDao = *internal.AdminMemberRoleDao

// adminMemberRoleDao is the data access object for table hg_admin_member_role.
// You can define custom methods on it to extend its functionality as you wish.
type adminMemberRoleDao struct {
	internalAdminMemberRoleDao
}

var (
	// AdminMemberRole is globally common accessible object for table hg_admin_member_role operations.
	AdminMemberRole = adminMemberRoleDao{
		internal.NewAdminMemberRoleDao(),
	}
)

// Fill with you ideas below.
