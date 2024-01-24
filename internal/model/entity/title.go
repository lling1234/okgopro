// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Title is the golang structure for table title.
type Title struct {
	Id         int64       `json:"id"         ` //
	Name       string      `json:"name"       ` // 内容标题
	SourceKey  string      `json:"sourceKey"  ` // 内容key
	Sort       string      `json:"sort"       ` // 排序
	Link       string      `json:"link"       ` // 链接
	Icon       string      `json:"icon"       ` // 图标
	Remark1    string      `json:"remark1"    ` // 备用1
	Remark2    string      `json:"remark2"    ` // 备用2
	Remark3    string      `json:"remark3"    ` // 备用3
	TenantId   string      `json:"tenantId"   ` // 租户id，可为空
	CreateUser string      `json:"createUser" ` // 创建人
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateUser string      `json:"updateUser" ` // 更新人
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	Status     int         `json:"status"     ` // 状态
	IsDeleted  int         `json:"isDeleted"  ` // 是否删除：0.正常,1.删除
}
