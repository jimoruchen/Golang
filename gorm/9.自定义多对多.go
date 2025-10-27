package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
)

func UserManyToMany() {
	//创建用户连带创建文章
	//global.DB.SetupJoinTable(&models.User1Model{}, "CollArticleList", &models.User2ArticleModel{})
	//err := global.DB.Create(&models.User1Model{
	//	Name: "李四",
	//	CollArticleList: []models.Article1Model{
	//		{Title: "Python"},
	//		{Title: "JS"},
	//	},
	//}).Error
	//err := global.DB.Create(&models.User1Model{
	//	Name: "xxx",
	//	CollArticleList: []models.Article1Model{
	//		{ID: 1},
	//		{ID: 2},
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//查某个用户他收藏的文章
	//var userID = 1
	//var userArticleList []models.User2ArticleModel
	//global.DB.Preload("UserModel").Preload("ArticleModel").Find(&userArticleList, "user_id = ?", userID)
	//fmt.Println(userArticleList)

	//type UserCollArticleResponse struct {
	//	Name         string
	//	UserID       int64
	//	ArticleID    int64
	//	ArticleTitle string
	//	Date         time.Time
	//}
	//var userID = 1
	//var userArticleList []models.User2ArticleModel
	//var collList []UserCollArticleResponse
	//global.DB.Preload("UserModel").Preload("ArticleModel").Find(&userArticleList, "user_id = ?", userID)
	//for _, userArticle := range userArticleList {
	//	collList = append(collList, UserCollArticleResponse{
	//		Name:         userArticle.UserModel.Name,
	//		UserID:       userArticle.UserID,
	//		ArticleID:    userArticle.ArticleID,
	//		ArticleTitle: userArticle.ArticleModel.Title,
	//		Date:         userArticle.CreatedAt,
	//	})
	//}
	//fmt.Println(collList)

	//用户取消收藏，清空关联关系
	var user models.User1Model
	global.DB.Take(&user, "id=?", 6)
	global.DB.Model(&user).Association("CollArticleList").Clear()
}

func main() {
	global.Connect()
	//global.Migrate()
	UserManyToMany()
}
