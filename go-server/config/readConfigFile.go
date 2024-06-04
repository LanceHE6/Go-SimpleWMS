package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// 定义你的数据结构

type Config struct {
	SERVER struct {
		PORT string `yaml:"port"`
		MODE string `yaml:"mode"`
		LOG  struct {
			PATH      string `yaml:"path"`
			MAX_FILES int    `yaml:"max_files"`
		} `yaml:"log"`
		SECRET_KEY string `yaml:"secretKey"`
	} `yaml:"server"`
	DB struct {
		MYSQL struct {
			HOST     string `yaml:"host"`
			PORT     string `yaml:"port"`
			ACCOUNT  string `yaml:"account"`
			PASSWORD string `yaml:"password"`
			DBNAME   string `yaml:"dbname"`
		} `yaml:"mysql"`
	} `yaml:"db"`
}

var ServerConfig Config

// 创建一个函数来读取和解析YAML文件

func init() {
	// 检查文件是否存在
	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		// 如果文件不存在，创建并写入默认值
		defaultConfig := Config{
			SERVER: struct {
				PORT string `yaml:"port"`
				MODE string `yaml:"mode"`
				LOG  struct {
					PATH      string `yaml:"path"`
					MAX_FILES int    `yaml:"max_files"`
				} `yaml:"log"`
				SECRET_KEY string `yaml:"secretKey"`
			}(struct {
				PORT string `yaml:"port"`
				MODE string `yaml:"mode"`
				LOG  struct {
					PATH      string `yaml:"path"`
					MAX_FILES int    `yaml:"max_file"`
				} `yaml:"log"`
				SECRET_KEY string `yaml:"secretKey"`
			}{
				PORT: "8080",
				MODE: "debug",
				LOG: struct {
					PATH      string `yaml:"path"`
					MAX_FILES int    `yaml:"max_file"`
				}{
					PATH:      "logs",
					MAX_FILES: 15,
				},
				SECRET_KEY: "simple_wms_secret_key",
			}),
			DB: struct {
				MYSQL struct {
					HOST     string `yaml:"host"`
					PORT     string `yaml:"port"`
					ACCOUNT  string `yaml:"account"`
					PASSWORD string `yaml:"password"`
					DBNAME   string `yaml:"dbname"`
				} `yaml:"mysql"`
			}{
				MYSQL: struct {
					HOST     string `yaml:"host"`
					PORT     string `yaml:"port"`
					ACCOUNT  string `yaml:"account"`
					PASSWORD string `yaml:"password"`
					DBNAME   string `yaml:"dbname"`
				}{
					HOST:     "localhost",
					PORT:     "3306",
					ACCOUNT:  "root",
					PASSWORD: "root",
					DBNAME:   "simple_wms",
				},
			},
		}

		defaultBytes, err := yaml.Marshal(&defaultConfig)
		if err != nil {
			_ = fmt.Errorf("can not marshal the default config")
			return
		}

		err = ioutil.WriteFile("config.yaml", defaultBytes, 0644)
		if err != nil {
			_ = fmt.Errorf("can not write the default config to file")
			return
		}
	}

	// 读取文件
	fileBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		_ = fmt.Errorf("can not read the config file")
		return
	}

	// 解析YAML文件
	err = yaml.Unmarshal(fileBytes, &ServerConfig)
	if err != nil {
		_ = fmt.Errorf("can not unmarshal the config file")
		return
	}

}
