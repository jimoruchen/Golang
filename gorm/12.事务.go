package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"

	"gorm.io/gorm"
)

func test1() {
	var zhangsan = models.UserModel{ID: 1}
	var lisi = models.UserModel{ID: 2}
	global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&zhangsan).Update("money", gorm.Expr("money - ?", 100)).Error
		if err != nil {
			return err
		}
		err = tx.Model(&lisi).Update("money", gorm.Expr("money + ?", 100)).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func test2() {
	var zhangsan = models.UserModel{ID: 1}
	var lisi = models.UserModel{ID: 2}
	tx := global.DB.Begin()
	err := tx.Model(&zhangsan).Update("money", gorm.Expr("money - ?", 100)).Error
	//err = errors.New("出错了")
	if err != nil {
		tx.Rollback()
	}
	err = tx.Model(&lisi).Update("money", gorm.Expr("money + ?", 100)).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func main() {
	global.Connect()
	//global.Migrate()
	test2()
}
