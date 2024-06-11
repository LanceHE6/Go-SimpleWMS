package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type File struct {
	Path string `json:"path"`
}
type FileList []File

func (fl *FileList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to cast value to []byte")
	}

	err := json.Unmarshal(bytes, fl)
	if err != nil {
		return err
	}

	return nil
}

func (fl FileList) Value() (driver.Value, error) {
	bytes, err := json.Marshal(fl)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

type Goods struct {
	Base
	Gid          string   `gorm:"primary_key;index" json:"gid"`
	GoodsCode    string   `gorm:"unique" json:"goods_code"`
	Name         string   `json:"name"`
	Model        string   `json:"model"`
	GoodsType    string   `gorm:"default:''" json:"goods_type"`
	Manufacturer string   `gorm:"default:''" json:"manufacturer"`
	Unit         string   `json:"unit"`
	Images       FileList `gorm:"type:json" json:"images"`
	Files        FileList `gorm:"type:json" json:"files"`
	UnitPrice    float64  `gorm:"default:0" json:"unit_price"`
}
