package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

var (
	db   *sqlx.DB
	once sync.Once
)

// DBConfig 用于映射YAML文件的内容
type DBConfig struct {
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

func InitDB() *sqlx.DB {
	once.Do(func() {
		var err error

		// 读取YAML文件
		data, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Fatalf("Error reading YAML file: %s\n", err)
		}

		// 创建一个DBConfig实例
		var config DBConfig

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

		// 格式为：user:password@tcp(localhost:5555)/dbname?charset=utf8&parseTime=True&loc=Local
		// 先连接到MySQL服务器，不指定数据库

		db, err = sqlx.Connect("mysql", dsn)
		if err != nil {
			log.Fatal("Connect DB error: " + err.Error())
		}

		// 创建数据库
		_, err = db.Exec("CREATE DATABASE IF NOT EXISTS simple_wms")
		if err != nil {
			log.Fatal("Create DB error: " + err.Error())
		}

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DB.MySQL.Account,
			config.DB.MySQL.Password,
			config.DB.MySQL.Host,
			config.DB.MySQL.Port,
			config.DB.MySQL.Database,
		)
		// 然后连接到新创建的数据库
		db, err = sqlx.Connect("mysql", dsn)
		if err != nil {
			log.Fatal("Connect DB error: " + err.Error())
		}
		log.Println("Check tables...")

		// 表创建语句
		sqlFile := "createTable.sql"
		sqlContent, err := ioutil.ReadFile(sqlFile)
		if err != nil {
			log.Fatal("Failed to read SQL file: " + err.Error())
			return
		}

		statements := strings.Split(string(sqlContent), ";")

		for _, statement := range statements {
			trimmedStatement := strings.TrimSpace(statement)
			if trimmedStatement != "" {
				_, err := db.Exec(trimmedStatement)
				if err != nil {
					log.Fatal("Failed to execute SQL statement: " + err.Error())
					return
				}
			}
		}

		log.Println("Check tables complete")
	})

	return db
}

func CloseDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Println("Close DB error: ", err)
		}
	}
}

func GetDbConnection() (*sql.Tx, error) {
	// 开始一个新的事务
	tx, err := db.Begin()
	if err != nil {
		log.Println("error: Cannot begin transaction")
		return nil, err
	}
	return tx, nil
}
