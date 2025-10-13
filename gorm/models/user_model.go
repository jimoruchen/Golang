package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18;check:age > 0"`
	CreatedAt time.Time
}

func (u UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")
	return nil
}
