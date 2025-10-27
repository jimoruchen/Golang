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
	UserDetailModel *UserDetailModel `gorm:"foreignkey:UserID"` //去 user_detail_models 表里找 UserID = UserModel.ID 的记录。
	//当我从 UserModel 查找它的 UserDetailModel 时，请使用 UserDetailModel 这个表中的 UserID 字段，去匹配当前 UserModel 的主键（ID）。
	CreatedAt time.Time
}

type UserDetailModel struct {
	ID        int64
	UserID    int64      `gorm:"unique"`            //一对一关系需要加上唯一约束
	UserModel *UserModel `gorm:"foreignKey:UserID"` //UserDetailModel 的 UserID 字段是外键，引用 UserModel.ID。
	// UserModel 这个关联对象，是通过当前表（user_detail_models）的 UserID 字段来查找的。
	Email string `gorm:"size:64"`
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
