// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package menu

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table sys_menu.
type Entity struct {
	MenuId     int64       `orm:"menu_id,primary" json:"menu_id"`     // 菜单ID
	MenuName   string      `orm:"menu_name"       json:"menu_name"`   // 菜单名称
	ParentId   int64       `orm:"parent_id"       json:"parent_id"`   // 父菜单ID
	OrderNum   int         `orm:"order_num"       json:"order_num"`   // 显示顺序
	Url        string      `orm:"url"             json:"url"`         // 请求地址
	Target     string      `orm:"target"          json:"target"`      // 打开方式（menuItem页签 menuBlank新窗口）
	MenuType   string      `orm:"menu_type"       json:"menu_type"`   // 菜单类型（M目录 C菜单 F按钮）
	Visible    string      `orm:"visible"         json:"visible"`     // 菜单状态（0显示 1隐藏）
	Perms      string      `orm:"perms"           json:"perms"`       // 权限标识
	Icon       string      `orm:"icon"            json:"icon"`        // 菜单图标
	CreateBy   string      `orm:"create_by"       json:"create_by"`   // 创建者
	CreateTime *gtime.Time `orm:"create_time"     json:"create_time"` // 创建时间
	UpdateBy   string      `orm:"update_by"       json:"update_by"`   // 更新者
	UpdateTime *gtime.Time `orm:"update_time"     json:"update_time"` // 更新时间
	Remark     string      `orm:"remark"          json:"remark"`      // 备注
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
