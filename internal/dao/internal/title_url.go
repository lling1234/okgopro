// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TitleUrlDao is the data access object for table title_url.
type TitleUrlDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TitleUrlColumns // columns contains all the column names of Table for convenient usage.
}

// TitleUrlColumns defines and stores column names for table title_url.
type TitleUrlColumns struct {
	Id         string //
	Title      string // 历史ID（即变更前ID）
	Extra      string // 供应商ID
	Link       string // 银行信息ID
	WebUrlId   string // 基本户户名
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

// titleUrlColumns holds the columns for table title_url.
var titleUrlColumns = TitleUrlColumns{
	Id:         "id",
	Title:      "title",
	Extra:      "extra",
	Link:       "link",
	WebUrlId:   "web_url_id",
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

// NewTitleUrlDao creates and returns a new DAO object for table data access.
func NewTitleUrlDao() *TitleUrlDao {
	return &TitleUrlDao{
		group:   "default",
		table:   "title_url",
		columns: titleUrlColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TitleUrlDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TitleUrlDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TitleUrlDao) Columns() TitleUrlColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TitleUrlDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TitleUrlDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TitleUrlDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
