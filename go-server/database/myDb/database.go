package myDb

import (
	"Go_simpleWMS/database/model"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var db *gorm.DB

// DbConfig 用于映射YAML文件的内容
type dbConfig struct {
	DB struct {
		MySQL struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Account  string `yaml:"account"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"mysql"`
	} `yaml:"db"`
}

func Init() {
	var err error
	// 读取YAML文件
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	// 创建一个DBConfig实例
	var config dbConfig

	// 解析YAML文件
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	// 创建MySQL连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		config.DB.MySQL.Account,
		config.DB.MySQL.Password,
		config.DB.MySQL.Host,
		config.DB.MySQL.Port,
	)

	// 连接到MySQL
	tdb, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	// 创建数据库
	_, err = tdb.Exec("CREATE DATABASE IF NOT EXISTS " + config.DB.MySQL.Database)
	if err != nil {
		return
	}
	// 关闭数据库连接
	err = tdb.Close()
	if err != nil {
		return
	}

	// 创建MySQL连接字符串

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB.MySQL.Account,
		config.DB.MySQL.Password,
		config.DB.MySQL.Host,
		config.DB.MySQL.Port,
		config.DB.MySQL.Database,
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Cannot connect to MySQL database: %v", err)
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