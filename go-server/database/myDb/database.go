package myDb

import (
	"Go_simpleWMS/config"
	"Go_simpleWMS/database/model"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var db *gorm.DB

func Init() {
	var err error

	// 创建MySQL连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		config.ServerConfig.DB.MYSQL.ACCOUNT,
		config.ServerConfig.DB.MYSQL.PASSWORD,
		config.ServerConfig.DB.MYSQL.HOST,
		config.ServerConfig.DB.MYSQL.PORT,
	)

	// 连接到MySQL
	tdb, err := sql.Open("mysql", dsn)
	if err != nil {
		_ = fmt.Errorf("can not connect to database")
		os.Exit(-1)
		return
	}

	// 创建数据库
	_, err = tdb.Exec("CREATE DATABASE IF NOT EXISTS " + config.ServerConfig.DB.MYSQL.DBNAME)
	if err != nil {
		_ = fmt.Errorf("can not create database")
		os.Exit(-2)
		return
	}
	// 关闭数据库连接
	err = tdb.Close()
	if err != nil {
		_ = fmt.Errorf("can not close the database")
		os.Exit(-3)
		return
	}

	// 创建MySQL连接字符串

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ServerConfig.DB.MYSQL.ACCOUNT,
		config.ServerConfig.DB.MYSQL.PASSWORD,
		config.ServerConfig.DB.MYSQL.HOST,
		config.ServerConfig.DB.MYSQL.PORT,
		config.ServerConfig.DB.MYSQL.DBNAME,
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Cannot connect to MYSQL database: %v", err)
		os.Exit(-4)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Department{})
	db.AutoMigrate(&model.Unit{})
	// 声明外键，级联删除和更新
	db.AutoMigrate(&model.Staff{}).AddForeignKey("department", "departments(did)", "SET NULL", "CASCADE")
	db.AutoMigrate(&model.InventoryType{})
	db.AutoMigrate(&model.Warehouse{}).AddForeignKey("manager", "staffs(sid)", "SET NULL", "CASCADE")
	db.AutoMigrate(&model.GoodsType{})
	db.AutoMigrate(&model.Goods{}).AddForeignKey("goods_type", "goods_types(gtid)", "SET NULL", "CASCADE").AddForeignKey("warehouse", "warehouses(wid)", "SET NULL", "CASCADE").AddForeignKey("unit", "units(unid)", "SET NULL", "CASCADE")

}

func CloseMyDb() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Println("Close Db error: ", err)
		}
	}
}

func GetMyDbConnection() *gorm.DB {
	return db
}
