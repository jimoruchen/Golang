package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID              int              `gorm:"primary_key"`
	Name            string           `gorm:"size:16;not null;unique"`
	Age             int              `gorm:"default:18;check:age > 0"`
	UserDetailModel *UserDetailModel `gorm:"foreignkey:UserID"`
	CreatedAt       time.Time
}

type UserDetailModel struct {
	ID        int64
	UserID    int64      `gorm:"unique"` //一对一关系需要加上唯一约束
	UserModel *UserModel `gorm:"foreignKey:UserID"`
	Email     string     `gorm:"size:64"`
}

func (u UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")
	return nil
}

func (u UserModel) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("更新的钩子函数")
	return nil
}

func (u UserModel) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("删除的钩子函数")
	return nil
}

func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("查询钩子")
	return
}
