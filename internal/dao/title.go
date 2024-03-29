// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"okgopro/internal/dao/internal"
)

// internalTitleDao is internal type for wrapping internal DAO implements.
type internalTitleDao = *internal.TitleDao

// titleDao is the data access object for table title.
// You can define custom methods on it to extend its functionality as you wish.
type titleDao struct {
	internalTitleDao
}

var (
	// Title is globally public accessible object for table title operations.
	Title = titleDao{
		internal.NewTitleDao(),
	}
)

// Fill with you ideas below.
