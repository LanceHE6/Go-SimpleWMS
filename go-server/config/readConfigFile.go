package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// 定义你的数据结构

type Config struct {
	SERVER struct {
		PORT string `yaml:"port"`
		MODE string `yaml:"mode"`
		LOG  struct {
			PATH string `yaml:"path"`
			FILE string `yaml:"file"`
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
