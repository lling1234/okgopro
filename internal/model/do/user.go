// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta        `orm:"table:user, do:true"`
	Id            interface{} //
	Username      interface{} // 用户名
	Password      interface{} // 密码
	Phone         interface{} // 手机号
	Name          interface{} // 姓名
	RealName      interface{} // 联系人姓名
	Sex           interface{} // 性别 1.男 2.女
	ProvinceId    interface{} // 地址省
	CityId        interface{} // 地址市
	CountyId      interface{} // 地址县
	AddressDetail interface{} // 详细地址
	Birthday      interface{} // 出生日期 yyyy-mm-dd
	Email         interface{} // 邮箱
	Remark1       interface{} // 备用1
	Remark2       interface{} // 备用2
	Remark3       interface{} // 备用3
	TenantId      interface{} // 租户id
	CreateUser    interface{} // 创建人
	CreateTime    *gtime.Time // 创建时间
	UpdateUser    interface{} // 更新人
	UpdateTime    *gtime.Time // 更新时间
	Status        interface{} // 状态 2-正常 3-黑名单
	IsAdmin       interface{} // 是否管理员：1-个人用户 2-供应商 3-供应商管理员
	IsDeleted     interface{} // 是否删除：0.正常,1.删除
}
