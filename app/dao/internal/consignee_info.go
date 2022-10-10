// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ConsigneeInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ConsigneeInfoDao struct {
	gmvc.M                       // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB               // DB is the raw underlying database management object.
	Table   string               // Table is the table name of the DAO.
	Columns consigneeInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// ConsigneeInfoColumns defines and stores column names for table consignee_info.
type consigneeInfoColumns struct {
	Id        string // 收货地址表
	UserId    string //
	IsDefault string // 默认地址1  非默认0
	Name      string //
	Phone     string //
	Province  string //
	City      string //
	Town      string // 县区
	Street    string // 街道乡镇
	Detail    string // 地址详情
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

func NewConsigneeInfoDao() *ConsigneeInfoDao {
	return &ConsigneeInfoDao{
		M:     g.DB("default").Model("consignee_info").Safe(),
		DB:    g.DB("default"),
		Table: "consignee_info",
		Columns: consigneeInfoColumns{
			Id:        "id",
			UserId:    "user_id",
			IsDefault: "is_default",
			Name:      "name",
			Phone:     "phone",
			Province:  "province",
			City:      "city",
			Town:      "town",
			Street:    "street",
			Detail:    "detail",
			CreatedAt: "created_at",
			UpdatedAt: "updated_at",
			DeletedAt: "deleted_at",
		},
	}
}