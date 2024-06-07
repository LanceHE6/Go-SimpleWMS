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

	// 从环境变量获取配置
	account := os.Getenv("MYSQL_ACCOUNT")
	if account == "" {
		account = config.ServerConfig.DB.MYSQL.ACCOUNT
	}
	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = config.ServerConfig.DB.MYSQL.PASSWORD
	}
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = config.ServerConfig.DB.MYSQL.HOST
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = config.ServerConfig.DB.MYSQL.PORT
	}
	dbname := os.Getenv("MYSQL_DBNAME")
	if dbname == "" {
		dbname = config.ServerConfig.DB.MYSQL.DBNAME
	}

	// 创建MySQL连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		account,
		password,
		host,
		port,
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
		account,
		password,
		host,
		port,
		dbname,
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
	db.AutoMigrate(&model.Inventory{}).AddForeignKey("inventory_type", "inventory_types(itid)", "SET NULL", "CASCADE").AddForeignKey("warehouse", "warehouses(wid)", "SET NULL", "CASCADE")
	db.AutoMigrate(&model.Stock{}).AddForeignKey("goods", "goods(gid)", "SET NULL", "CASCADE").AddForeignKey("warehouse", "warehouses(wid)", "SET NULL", "CASCADE")

	// 初始化数据
	fmt.Println("Init data...")
	InitData()
	fmt.Println("Init data done.")
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
