// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id            int64       `json:"id"            ` //
	Username      string      `json:"username"      ` // 用户名
	Password      string      `json:"password"      ` // 密码
	Phone         string      `json:"phone"         ` // 手机号
	Name          string      `json:"name"          ` // 姓名
	RealName      string      `json:"realName"      ` // 联系人姓名
	Sex           int         `json:"sex"           ` // 性别 1.男 2.女
	ProvinceId    int         `json:"provinceId"    ` // 地址省
	CityId        int         `json:"cityId"        ` // 地址市
	CountyId      int         `json:"countyId"      ` // 地址县
	AddressDetail string      `json:"addressDetail" ` // 详细地址
	Birthday      string      `json:"birthday"      ` // 出生日期 yyyy-mm-dd
	Email         string      `json:"email"         ` // 邮箱
	Remark1       string      `json:"remark1"       ` // 备用1
	Remark2       string      `json:"remark2"       ` // 备用2
	Remark3       string      `json:"remark3"       ` // 备用3
	TenantId      string      `json:"tenantId"      ` // 租户id
	CreateUser    int64       `json:"createUser"    ` // 创建人
	CreateTime    *gtime.Time `json:"createTime"    ` // 创建时间
	UpdateUser    int64       `json:"updateUser"    ` // 更新人
	UpdateTime    *gtime.Time `json:"updateTime"    ` // 更新时间
	Status        int         `json:"status"        ` // 状态 2-正常 3-黑名单
	IsAdmin       int         `json:"isAdmin"       ` // 是否管理员：1-个人用户 2-供应商 3-供应商管理员
	IsDeleted     int         `json:"isDeleted"     ` // 是否删除：0.正常,1.删除
}
