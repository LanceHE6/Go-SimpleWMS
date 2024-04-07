package myDb

import (
	"Go_simpleWMS/database/model"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
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
	db.AutoMigrate(&model.Staff{}).AddForeignKey("did", "departments(did)", "SET NULL", "CASCADE")
	db.AutoMigrate(&model.InventoryType{})
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
