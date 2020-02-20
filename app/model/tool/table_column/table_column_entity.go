// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package table_column

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table gen_table_column.
type Entity struct {
	ColumnId      int64       `orm:"column_id,primary" json:"column_id"`      // 编号
	TableId       int64       `orm:"table_id"          json:"table_id"`       // 归属表编号
	ColumnName    string      `orm:"column_name"       json:"column_name"`    // 列名称
	ColumnComment string      `orm:"column_comment"    json:"column_comment"` // 列描述
	ColumnType    string      `orm:"column_type"       json:"column_type"`    // 列类型
	GoType        string      `orm:"go_type"         json:"go_type"`          // GO类型
	GoField       string      `orm:"go_field"        json:"go_field"`         // GO字段名
	HtmlField     string      `orm:"html_field"        json:"html_field"`     // html字段名
	IsPk          string      `orm:"is_pk"             json:"is_pk"`          // 是否主键（1是）
	IsIncrement   string      `orm:"is_increment"      json:"is_increment"`   // 是否自增（1是）
	IsRequired    string      `orm:"is_required"       json:"is_required"`    // 是否必填（1是）
	IsInsert      string      `orm:"is_insert"         json:"is_insert"`      // 是否为插入字段（1是）
	IsEdit        string      `orm:"is_edit"           json:"is_edit"`        // 是否编辑字段（1是）
	IsList        string      `orm:"is_list"           json:"is_list"`        // 是否列表字段（1是）
	IsQuery       string      `orm:"is_query"          json:"is_query"`       // 是否查询字段（1是）
	QueryType     string      `orm:"query_type"        json:"query_type"`     // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      string      `orm:"html_type"         json:"html_type"`      // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      string      `orm:"dict_type"         json:"dict_type"`      // 字典类型
	Sort          int         `orm:"sort"              json:"sort"`           // 排序
	CreateBy      string      `orm:"create_by"         json:"create_by"`      // 创建者
	CreateTime    *gtime.Time `orm:"create_time"       json:"create_time"`    // 创建时间
	UpdateBy      string      `orm:"update_by"         json:"update_by"`      // 更新者
	UpdateTime    *gtime.Time `orm:"update_time"       json:"update_time"`    // 更新时间
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
