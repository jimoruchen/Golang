package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
)

func oneToMany() {
	//创建班级连带创建多个学生
	//err := global.DB.Create(&models.ClassModel{
	//	Name: "class1",
	//	StudentList: []models.StudentModel{
	//		{Name: "student1"},
	//		{Name: "student2"},
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//创建学生自带班级
	//global.DB.Create(&models.StudentModel{
	//	Name: "student3",
	//	ClassModel: &models.ClassModel{
	//		Name: "class2",
	//	},
	//})

	//创建学生关联已有班级
	//var class models.ClassModel
	//global.DB.Take(&class, 1)
	//global.DB.Create(&models.StudentModel{
	//	Name:       "student4",
	//	ClassModel: &class,
	//})
	//global.DB.Create(&models.StudentModel{
	//	Name:       "student6",
	//	ClassModel: &models.ClassModel{ID: 2},
	//})

	//查询
	//var class models.ClassModel
	//global.DB.Preload("StudentList").Find(&class, "name = ?", "class1")
	//fmt.Println(class.Name, class.StudentList, len(class.StudentList))

	//带条件查询
	//var class models.ClassModel
	//global.DB.Preload("StudentList", "name = ?", "student1").Find(&class, "name = ?", "class1")
	//fmt.Println(class.Name, class.StudentList, len(class.StudentList))

	//查班级的学生列表和总数
	//var class models.ClassModel
	//global.DB.Take(&class, "name = ?", "class1")
	//var studentList []models.StudentModel
	//global.DB.Model(&class).Association("StudentList").Find(&studentList)
	//count := global.DB.Model(&class).Association("StudentList").Count()
	//fmt.Println(studentList, count)

	//一个学生退出班级
	//var class models.ClassModel
	//global.DB.Find(&class, 1)
	//global.DB.Model(&class).Association("StudentList").Delete([]models.StudentModel{{ID: 5}})

	//var class models.ClassModel
	//global.DB.Find(&class, 1)
	//var student models.StudentModel
	//global.DB.First(&student, 5)
	//global.DB.Model(&class).Association("StudentList").Delete([]models.StudentModel{student})

	//班级学生全部退出
	//var class models.ClassModel
	//global.DB.Find(&class, 2)
	//global.DB.Model(&class).Association("StudentList").Clear()

	//班级学生全部换成其他学生
	//var student7 = models.StudentModel{
	//	ID:   7,
	//	Name: "student7",
	//}
	//var class models.ClassModel
	//global.DB.Find(&class, 2)
	//global.DB.Model(&class).Association("StudentList").Replace([]models.StudentModel{student7})

	//已退出学生又加入班级
	var class models.ClassModel
	global.DB.Find(&class, 2)
	global.DB.Model(&class).Association("StudentList").Append([]models.StudentModel{{ID: 3}, {ID: 6}})
}

func main() {
	global.Connect()
	//global.Migrate()
	oneToMany()
}
