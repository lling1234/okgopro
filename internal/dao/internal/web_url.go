// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebUrlDao is the data access object for table web_url.
type WebUrlDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns WebUrlColumns // columns contains all the column names of Table for convenient usage.
}

// WebUrlColumns defines and stores column names for table web_url.
type WebUrlColumns struct {
	Id         string //
	Name       string // 历史ID（即变更前ID）
	SourceKey  string // 基本户户名
	Sort       string // 供应商ID
	Link       string // 银行信息ID
	Icon       string // 基本户户名
	Remark1    string // 备用1
	Remark2    string // 备用2
	Remark3    string // 备用3
	TenantId   string // 租户id，可为空
	CreateUser string // 创建人
	CreateTime string // 创建时间
	UpdateUser string // 更新人
	UpdateTime string // 更新时间
	Status     string // 状态
	IsDeleted  string // 是否删除：0.正常,1.删除
}

// webUrlColumns holds the columns for table web_url.
var webUrlColumns = WebUrlColumns{
	Id:         "id",
	Name:       "name",
	SourceKey:  "source_key",
	Sort:       "sort",
	Link:       "link",
	Icon:       "icon",
	Remark1:    "remark1",
	Remark2:    "remark2",
	Remark3:    "remark3",
	TenantId:   "tenant_id",
	CreateUser: "create_user",
	CreateTime: "create_time",
	UpdateUser: "update_user",
	UpdateTime: "update_time",
	Status:     "status",
	IsDeleted:  "is_deleted",
}

// NewWebUrlDao creates and returns a new DAO object for table data access.
func NewWebUrlDao() *WebUrlDao {
	return &WebUrlDao{
		group:   "default",
		table:   "web_url",
		columns: webUrlColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WebUrlDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WebUrlDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WebUrlDao) Columns() WebUrlColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WebUrlDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WebUrlDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WebUrlDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
