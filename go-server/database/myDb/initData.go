package myDb

import (
	"Go_simpleWMS/database/model"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func InitData() {
	db := GetMyDbConnection()
	// 初始化数据
	// 初始化出入库类型
	createIfNotExists(db, &model.InventoryType{
		Itid:     "_default1_",
		Name:     "默认入库",
		TypeCode: "default_in",
		Type:     1,
	}, "_default1_", "itid")
	createIfNotExists(db, &model.InventoryType{
		Itid:     "_default2_",
		Name:     "默认出库",
		TypeCode: "default_out",
		Type:     2,
	}, "_default2_", "itid")
	// 初始化部门
	createIfNotExists(db, &model.Department{
		Did:  "_default_",
		Name: "未知",
	}, "_default_", "did")
	// 初始化员工
	createIfNotExists(db, &model.Staff{
		Sid:        "_default_",
		Name:       "管理员",
		Department: "_default_", // 默认部门
	}, "_default_", "sid")
	// 初始化仓库
	createIfNotExists(db, &model.Warehouse{
		Wid:     "_default_",
		Name:    "默认仓库",
		Manager: "_default_", // 默认管理员
	}, "_default_", "wid")
	// 初始化货品类型
	createIfNotExists(db, &model.GoodsType{
		Gtid:     "_default_",
		Name:     "未分类",
		TypeCode: "default_type",
	}, "_default_", "gtid")
	// 初始化用户
	createIfNotExists(db, &model.User{
		Uid:        "u00000001",
		Account:    "admin",
		Password:   "123456",
		Permission: 3,
		Nickname:   "admin",
	}, "u00000001", "uid")
	// 初始化计量单位
	createIfNotExists(db, &model.Unit{
		Unid: "_default_",
		Name: "未知",
	}, "_default_", "unid")
}

func createIfNotExists(db *gorm.DB, value interface{}, id interface{}, idFieldName string) {
	// 检查数据是否存在
	if err := db.Where(idFieldName+" = ?", id).First(value).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 数据不存在，插入新数据
			if err := db.Create(value).Error; err != nil {
				// 插入错误
				log.Println(err)
				os.Exit(-100)
			}
		} else {
			// 查询错误
			log.Println(err)
			os.Exit(-200)
		}
	}
}
