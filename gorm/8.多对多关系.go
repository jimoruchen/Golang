package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
)

func ManyToMany() {
	//创建一篇文章，新增Tag
	//err := global.DB.Create(&models.ArticleModel{
	//	Title: "AAA",
	//	TagList: []models.TagModel{
	//		{Title: "111"},
	//		{Title: "222"},
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//创建一篇文章，选择Tag
	//var TagList []models.TagModel
	//TagIDList := []int{1, 2}
	//global.DB.Find(&TagList, "id in ?", TagIDList)
	//global.DB.Create(&models.ArticleModel{
	//	Title:   "BBB",
	//	TagList: TagList,
	//})

	//查文章时把对应的标签带出来
	//var article models.ArticleModel
	//global.DB.Preload("TagList").Find(&article, 1)
	//fmt.Println(article)

	//将文章2的标签更新为1和2
	var article models.ArticleModel
	global.DB.Find(&article, 2)
	global.DB.Model(&article).Association("TagList").Replace([]models.TagModel{
		{ID: 1},
		{ID: 2},
	})
	//global.DB.Model(&article).Association("TagList").Append([]models.TagModel{
	//	{ID: 2},
	//})
}

func main() {
	global.Connect()
	//global.Migrate()
	ManyToMany()
}
