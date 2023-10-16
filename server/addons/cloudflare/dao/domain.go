// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/addons/cloudflare/dao/internal"
)

// internalDomainDao is internal type for wrapping internal DAO implements.
type internalDomainDao = *internal.DomainDao

// domainDao is the data access object for table ue_cf_domain.
// You can define custom methods on it to extend its functionality as you wish.
type domainDao struct {
	internalDomainDao
}

var (
	// Domain is globally public accessible object for table ue_cf_domain operations.
	Domain = domainDao{
		internal.NewDomainDao(),
	}
)

// Fill with you ideas below.
