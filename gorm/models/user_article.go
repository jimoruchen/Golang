package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User1Model struct {
	ID              int64
	Name            string
	CollArticleList []Article1Model `gorm:"many2many:user2_article_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
}
type Article1Model struct {
	ID    int64
	Title string `gorm:"size:32"`
}
type User2ArticleModel struct {
	UserID       int64         `gorm:"primaryKey"`
	UserModel    User1Model    `gorm:"foreignKey:UserID"` //使用 UserID 字段作为外键，关联到 User1Model.ID
	ArticleID    int64         `gorm:"primaryKey"`
	ArticleModel Article1Model `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time     `json:"createdAt"`
	Title        string        `gorm:"size:32" json:"title"`
}

func (u *User2ArticleModel) BeforeCreate(tx *gorm.DB) error {
	var articleTitle string
	tx.Model(&Article1Model{}).Where("id = ?", u.ArticleID).Select("title").Scan(&articleTitle)
	u.Title = articleTitle
	fmt.Println("创建的钩子函数", u.ArticleID, u.Title)
	return nil
}

func (User2ArticleModel) TableName() string {
	return "user2_article_models"
}
