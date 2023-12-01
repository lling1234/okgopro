// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TitleUrl is the golang structure for table title_url.
type TitleUrl struct {
	Id         int64       `json:"id"         ` //
	Title      string      `json:"title"      ` // 历史ID（即变更前ID）
	Extra      string      `json:"extra"      ` // 供应商ID
	Link       string      `json:"link"       ` // 银行信息ID
	WebUrlId   string      `json:"webUrlId"   ` // 基本户户名
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
