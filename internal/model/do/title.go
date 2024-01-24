// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Title is the golang structure of table title for DAO operations like Where/Data.
type Title struct {
	g.Meta     `orm:"table:title, do:true"`
	Id         interface{} //
	Name       interface{} // 内容标题
	SourceKey  interface{} // 内容key
	Sort       interface{} // 排序
	Link       interface{} // 链接
	Icon       interface{} // 图标
	Remark1    interface{} // 备用1
	Remark2    interface{} // 备用2
	Remark3    interface{} // 备用3
	TenantId   interface{} // 租户id，可为空
	CreateUser interface{} // 创建人
	CreateTime *gtime.Time // 创建时间
	UpdateUser interface{} // 更新人
	UpdateTime *gtime.Time // 更新时间
	Status     interface{} // 状态
	IsDeleted  interface{} // 是否删除：0.正常,1.删除
}
