// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id            string //
	Username      string // 用户名
	Password      string // 密码
	Phone         string // 手机号
	Name          string // 姓名
	RealName      string // 联系人姓名
	Sex           string // 性别 1.男 2.女
	ProvinceId    string // 地址省
	CityId        string // 地址市
	CountyId      string // 地址县
	AddressDetail string // 详细地址
	Birthday      string // 出生日期 yyyy-mm-dd
	Email         string // 邮箱
	Remark1       string // 备用1
	Remark2       string // 备用2
	Remark3       string // 备用3
	TenantId      string // 租户id
	CreateUser    string // 创建人
	CreateTime    string // 创建时间
	UpdateUser    string // 更新人
	UpdateTime    string // 更新时间
	Status        string // 状态 2-正常 3-黑名单
	IsAdmin       string // 是否管理员：1-个人用户 2-管理员
	IsDeleted     string // 是否删除：0.正常,1.删除
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:            "id",
	Username:      "username",
	Password:      "password",
	Phone:         "phone",
	Name:          "name",
	RealName:      "real_name",
	Sex:           "sex",
	ProvinceId:    "province_id",
	CityId:        "city_id",
	CountyId:      "county_id",
	AddressDetail: "address_detail",
	Birthday:      "birthday",
	Email:         "email",
	Remark1:       "remark1",
	Remark2:       "remark2",
	Remark3:       "remark3",
	TenantId:      "tenant_id",
	CreateUser:    "create_user",
	CreateTime:    "create_time",
	UpdateUser:    "update_user",
	UpdateTime:    "update_time",
	Status:        "status",
	IsAdmin:       "is_admin",
	IsDeleted:     "is_deleted",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
