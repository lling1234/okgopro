// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContentDao is the data access object for table content.
type ContentDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ContentColumns // columns contains all the column names of Table for convenient usage.
}

// ContentColumns defines and stores column names for table content.
type ContentColumns struct {
	Id         string //
	Title      string // 标题
	Extra      string // 拓展
	PrefixUrl  string // url前缀
	Url        string // url
	Link       string // 链接
	WebUrlId   string // 链接id
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

// contentColumns holds the columns for table content.
var contentColumns = ContentColumns{
	Id:         "id",
	Title:      "title",
	Extra:      "extra",
	PrefixUrl:  "prefix_url",
	Url:        "url",
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

// NewContentDao creates and returns a new DAO object for table data access.
func NewContentDao() *ContentDao {
	return &ContentDao{
		group:   "default",
		table:   "content",
		columns: contentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContentDao) Columns() ContentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
