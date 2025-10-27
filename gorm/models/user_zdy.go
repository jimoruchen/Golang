package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type UserInfo struct {
	Likes []string `json:"likes"`
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Json
func (j *UserInfo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := UserInfo{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j UserInfo) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type Card struct {
	Card []string `json:"card"`
}

type UserZdy struct {
	ID       uint
	Name     string   `gorm:"size:32"`
	Info     UserInfo `gorm:"type:longtext" json:"info"`
	CardInfo Card     `gorm:"type:longtext;serializer:json" json:"card_info"`
}
