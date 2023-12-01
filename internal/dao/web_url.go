// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"okgopro/internal/dao/internal"
)

// internalWebUrlDao is internal type for wrapping internal DAO implements.
type internalWebUrlDao = *internal.WebUrlDao

// webUrlDao is the data access object for table web_url.
// You can define custom methods on it to extend its functionality as you wish.
type webUrlDao struct {
	internalWebUrlDao
}

var (
	// WebUrl is globally public accessible object for table web_url operations.
	WebUrl = webUrlDao{
		internal.NewWebUrlDao(),
	}
)

// Fill with you ideas below.