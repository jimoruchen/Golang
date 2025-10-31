# Golang

## gorm

### 传统sql的增删改查

#### 数据库操作
```sql
# 查询全部的数据库
show databases;

# 创建数据库 并指定字符集和排序方式
create database gorm_new_db
    character set utf8mb4
    collate utf8mb4_unicode_ci;

# 选中数据库
use gorm_new_db;

# 删除数据库
drop database gorm_new_db;
```

#### 表操作
```sql
# 查询当前库的所有表
show tables;


# 建表
create table users
(
    id         bigint unsigned auto_increment primary key,
    name       varchar(50)  not null default '',
    age        int unsigned not null default 0,
    email      varchar(100) not null default '',
    created_at datetime              default current_timestamp,
    updated_at datetime              default current_timestamp on update current_timestamp
) engine = innodb
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci;

# 插入数据
insert into users (name, age, email)
values ('张三', 18, 'zhangsan@example.com'),
       ('李四', 20, 'lisi@example.com');

# 查询数据
select  * from users;


# 删除数据
delete from users where id = 1;

# 更新数据
update users set name = '张三丰', age = 38 where id = 2;

# 删除表
drop table users;
```

### 通过go去操作sql

#### 安装驱动
```shell
go get -u github.com/go-sql-driver/mysql
```

#### 执行sql
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:200088@tcp(127.0.0.1:3306)/gorm_db?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败 %s", err)
	}
	fmt.Println(db)

	//插入
	//res, err := db.Exec("insert into users (name, age, email) values ('张三', 18, '111@qq.com')")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res)

	//查询
	rows, _ := db.Query("select id, name, age from users")
	defer rows.Close()
	for rows.Next() { //用于逐行推进结果集的游标
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age) //将当前行的列值扫描（复制）到 Go 变量中。
		fmt.Println(id, name, age)
	}

	var id1 int
	var name1 string
	var age1 int
	db.QueryRow("select id, name, age from users").Scan(&id1, &name1, &age1) //查一条
	fmt.Println(id1, name1, age1)
}
```

### gorm引入

```go
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	dsn := "root:200088@tcp(127.0.0.1:3306)/gorm_db?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db)

	var userList []User
	db.Find(&userList)
	fmt.Println(userList)

	db.Where("id < ?", 2).Find(&userList)
	fmt.Println(userList)

	db.Where("name = ?", "李四").Find(&userList)
	fmt.Println(userList)
}
```

### 单表模型与单表操作

#### 单表模型
```go
package main

import (
	"Golang/gorm/global"
	"fmt"
	"time"
)

type UserModel struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18"`
	CreatedAt time.Time
}

func migrate() {
	err := global.DB.AutoMigrate(&UserModel{})
	if err != nil {
		fmt.Println("表结构迁移失败", err)
		return
	}
	fmt.Println("表结构迁移成功")
}

func main() {
	global.Connect()
	migrate()
}
```

#### 单表操作
```go
package global

import (
	"Golang/gorm/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	err := DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatalf("数据库迁移失败 %s", err)
	}
	fmt.Printf("数据库迁移成功")
}

func Connect() {
	dst := "root:200088@tcp(127.0.0.1:3306)/gorm_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}

```
```go
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

func (u UserModel) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("更新的钩子函数")
	return nil
}

func (u UserModel) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("删除的钩子函数")
	return nil
}

```
```go
package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func insert() {
	//err := global.DB.Create(&models.UserModel{
	//	Name: "王五",
	//	Age:  18,
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//回填式创建
	//user := models.UserModel{
	//	Name: "李四",
	//	Age:  18,
	//}
	//err := global.DB.Create(&user).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(user.Name, user.ID, user.Age, user.CreatedAt)

	//批量插入
	//userList := []models.UserModel{
	//	{Name: "Alex", Age: 18},
	//	{Name: "Tom", Age: 20},
	//	{Name: "Invalid", Age: -5}, // 假设这个会触发数据库约束失败
	//}
	//for _, user := range userList {
	//	err := global.DB.Create(&user).Error
	//	if err != nil {
	//		fmt.Printf("创建用户 %s 失败：%v", user.Name, err)
	//		continue
	//	}
	//	fmt.Printf(" 创建用户 %s 成功", user.Name)
	//}

	//钩子函数
	err := global.DB.Create(&models.UserModel{
		Name: "王五1",
		Age:  18,
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}

func query() {
	var userList []models.UserModel

	//全部查询
	//err := global.DB.Find(&userList).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(userList)

	//特定查询
	global.DB.Where("age >= ?", 20).Find(&userList)
	//global.DB.Find(&userList, "age >= ?", 20)
	fmt.Println(userList)

	var user models.UserModel
	//查一个
	//global.DB.Take(&user)

	//根据主键去查
	//global.DB.Take(&user, 2)

	//根据主键排序查第一个
	//global.DB.First(&user)

	//根据主键排序查最后一个，.Debug()会打印实际的SQL
	//global.DB.Debug().Last(&user)

	//user.ID = 3 //查询会携带主键
	//global.DB.Take(&user)

	//查一条没查到会抛出错误没查到
	//err := global.DB.Take(&user, 111).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	fmt.Println("记录不存在")
	//	return
	//}

	//使用Limit(1).Find()，不会抛出错误没查到
	err := global.DB.Limit(1).Find(&user, 111).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("记录不存在")
		return
	}
	fmt.Println(user)
}

func save() {
	global.DB = global.DB.Debug()
	var user = models.UserModel{
		ID:        2,
		Name:      "", //可以更新零值
		Age:       20,
		CreatedAt: time.Now(),
	}
	global.DB.Save(&user)
}

//Save，有主键记录就是更新，并且可以更新零值，否则就是创建

func update() {
	var user = models.UserModel{ID: 1}

	global.DB.Model(&user).
		Where("id=?", 1).
		Update("age", 21)

	fmt.Println(user)
}

func updateColumn() {
	global.DB = global.DB.Debug()
	var user = models.UserModel{ID: 1}
	global.DB.Model(&user).
		Where("id=?", 1).
		UpdateColumn("age", 22).
		UpdateColumn("name", "张三")
}

func updates() {
	var user = models.UserModel{ID: 2}

	//global.DB.Model(&user).
	//	Updates(models.UserModel{
	//		Name: "张三",
	//		Age:  22,
	//	})

	global.DB.Model(&user).
		Where("id = ?", 2).
		Updates(map[string]interface{}{
			"age": 23,
		})

	fmt.Println(user)
}

func Delete() {
	global.DB = global.DB.Debug()

	//var user = models.UserModel{ID: 10}
	//global.DB.Delete(&user)

	//global.DB.Delete(&models.UserModel{}, 9)

	//global.DB.Delete(&models.UserModel{}, "name = ?", "Tom")

	//批量删除
	global.DB.Delete(&models.UserModel{}, []int{5, 7})
}

func main() {
	global.Connect()
	//save()
	//update()
	//updateColumn()
	//updates()
	Delete()
}

```
<img src="https://s2.loli.net/2025/10/14/MIAZbvuaf2V1zxB.png" >

### 单表高级查询
```go
package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"

	"gorm.io/gorm"
)

func highQuery() {
	global.DB = global.DB.Debug()

	//批量插入
	//userList := []models.UserModel{
	//	{Name: "张三", Age: 20},
	//	{Name: "李四", Age: 20},
	//	{Name: "王五", Age: 18},
	//	{Name: "Alex", Age: 18},
	//	{Name: "Bob", Age: 21},
	//}
	//for _, user := range userList {
	//	err := global.DB.Create(&user).Error
	//	if err != nil {
	//		fmt.Printf("create user err:%v\n", err)
	//		continue
	//	}
	//	fmt.Printf("成功，user:%#v\n", user.Name)
	//}

	//var user models.UserModel
	//global.DB.Where("age > ?", 18).Take(&user)
	//fmt.Println(user)

	//结构体
	//var user models.UserModel
	//global.DB.Where(models.UserModel{
	//	Name: "张三",
	//}).Take(&user)
	//fmt.Println(user)

	//map
	//var user models.UserModel
	//global.DB.Where(map[string]any{
	//	"name": "张三",
	//}).Take(&user)
	//fmt.Println(user)

	//嵌套
	//var user models.UserModel
	//query := global.DB.Where("age > 18")
	//global.DB.Where(query).Take(&user)
	//fmt.Println(user)

	//Or
	//var user []models.UserModel
	//global.DB.Or("name = ?", "张三").Or("age = ?", 18).Find(&user)
	//fmt.Println(user)

	//Not
	//var user []models.UserModel
	//global.DB.Not("age = ?", 18).Find(&user)
	//fmt.Println(user)

	//Order
	//var user []models.UserModel
	//global.DB.Order("age desc").Order("id asc").Find(&user)
	//fmt.Println(user)

	//Scan 只需要特定字段
	//var nameList []string
	//var user models.UserModel
	//global.DB.Model(user).Select("name").Scan(&nameList)
	//global.DB.Model(models.UserModel{}).Select("name").Scan(&nameList)
	//global.DB.Model(models.UserModel{}).Pluck("name", &nameList)
	//fmt.Println(nameList)

	//只需要几个特定的字段
	//type User struct {
	//	Name string
	//	Age  int
	//}
	//type User struct {
	//	Label string `gorm:"column:name"` //首字母必须大写
	//	Value int    `gorm:"column:age"`
	//}
	//
	//var userList []User
	//global.DB.Model(models.UserModel{}).Scan(&userList)
	//fmt.Println(userList)

	//分组
	//type User struct {
	//	Age   int
	//	Count int
	//}
	//var user []User
	//global.DB.Model(models.UserModel{}).
	//	Group("age").
	//	Select("age, count(*) as count").
	//	Scan(&user)
	//fmt.Println(user)		//[{22 1} {23 1} {18 2} {17 1}]

	//去重
	//var ageList []int
	//global.DB.Model(models.UserModel{}).Distinct("age").Pluck("age", &ageList)
	//fmt.Println(ageList)

	//分页
	//PaginateUsers(global.DB, 3, 2)

	//Scope
	//var users []models.UserModel
	//global.DB.Scopes(Age18).Find(&users)
	//fmt.Println(users)

	//var users []models.UserModel
	//global.DB.Scopes(NameIn("张三", "李四")).Find(&users)
	//fmt.Println(users)

	//原生SQL
	type User struct {
		Name string
		Age  int
	}
	var user []User
	global.DB.Raw("select name, age from user_models").Scan(&user)
	fmt.Println(user)
	global.DB.Exec("update user_models set age = 22 where id = 1")
}

func NameIn(NameList ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name in (?)", NameList)
	}
}

func Age18(tx *gorm.DB) *gorm.DB {
	return tx.Where("age > ?", 18)
}

func PaginateUsers(db *gorm.DB, pageNum, pageSize int) {
	var users []models.UserModel
	var count int64
	db.Model(&models.UserModel{}).Count(&count)
	offset := (pageNum - 1) * pageSize
	err := db.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("count: %d\n", count)
	fmt.Printf("%#v\n", users)
}

func main() {
	global.Connect()
	highQuery()
}
```

### 1对1关系
```go
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
	Money           int              `json:"money"`
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
```
```go
package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
)

func oneToOne() {
	//创建用户，连带着创建用户详情
	//err := global.DB.Create(&models.UserModel{
	//	Name: "张三",
	//	Age:  18,
	//	UserDetailModel: &models.UserDetailModel{
	//		Email: "111@qq.com",
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//创建用户详情，关联用户
	//err := global.DB.Create(&models.UserDetailModel{
	//	Email: "222@qq.com",
	//	//UserID: 13,
	//	UserModel: &models.UserModel{ID: 13},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//通过用户详情查用户		正查
	//var id = 12
	//var detail models.UserDetailModel
	//global.DB.Preload("UserModel").Take(&detail, "user_id = ?", id)
	//fmt.Println(detail.Email, detail.UserModel.Name)

	//反查
	//var id = 12
	//var user models.UserModel
	//global.DB.Preload("UserDetailModel").Take(&user, id)
	//fmt.Println(user, user.UserDetailModel.Email)

	//级联删除
	//var user models.UserModel
	//global.DB.Find(&user, 12)
	//fmt.Println(user)
	//global.DB.Select("UserDetailModel").Delete(&user)

	//set null
	var user models.UserModel
	global.DB.Find(&user, 13)
	global.DB.Model(&user).Association("UserDetailModel").Clear()
	global.DB.Delete(&user)
}

func main() {
	global.Connect()
	//global.Migrate()
	oneToOne()
}
```

### 1对多关系
```go
package models

type ClassModel struct {
	ID          int
	Name        string         `gorm:"size:12"`
	StudentList []StudentModel `gorm:"foreignkey:ClassID"`
}

type StudentModel struct {
	ID         int
	Name       string `gorm:"size:12"`
	ClassID    int
	ClassModel *ClassModel `gorm:"foreignkey:ClassID"`
}
```
```go
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
```

### 多对多关系
```go
package models

type ArticleModel struct {
	ID      int
	Title   string     `gorm:"size:32"`
	TagList []TagModel `gorm:"many2many:article_tags;"`
}

type TagModel struct {
	ID          int
	Title       string         `gorm:"size:32"`
	ArticleList []ArticleModel `gorm:"many2many:article_tags;"`
}
```
```go
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
```

### 自定义多对多关系
```go
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
```
```go
package main

import (
	"Golang/gorm/global"
)

func UserManyToMany() {
	//创建用户连带创建文章
	//global.DB.SetupJoinTable(&models.User1Model{}, "CollArticleList", &models.User2ArticleModel{})
	//err := global.DB.Create(&models.User1Model{
	//	Name: "张三",
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
	//var user models.User1Model
	//global.DB.Take(&user, "id=?", 6)
	//global.DB.Model(&user).Association("CollArticleList").Clear()
}

func main() {
	global.Connect()
	//global.Migrate()
	UserManyToMany()
}
```

### 自定义数据类型
```go
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
```
```go
package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"fmt"
)

func UserZdy() {
	//创建
	//err := global.DB.Create(&models.UserZdy{
	//	Name: "张三",
	//	Info: models.UserInfo{
	//		Likes: []string{"唱", "跳", "rap"},
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//查询
	//var userZdy models.UserZdy
	//global.DB.Take(&userZdy)
	//fmt.Println(userZdy, userZdy.Info.Likes)

	//创建
	//err := global.DB.Create(&models.UserZdy{
	//	Name: "张三",
	//	Info: models.UserInfo{
	//		Likes: []string{"唱", "跳", "rap"},
	//	},
	//	CardInfo: models.Card{
	//		Card: []string{"法拉利", "兰博基尼"},
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//查询
	var userZdy models.UserZdy
	global.DB.Last(&userZdy)
	fmt.Println(userZdy, userZdy.Info.Likes, userZdy.CardInfo.Card)
}

func main() {
	global.Connect()
	//global.Migrate()
	UserZdy()
}
```

### 枚举类型
```go
package main

import (
	"Golang/gorm/global"
	"Golang/gorm/models"
	"encoding/json"
	"fmt"
)

func test() {
	//global.DB.Create(&models.LogModel{
	//	Title: "用户注册",
	//	Level: models.InfoLevel,
	//})

	var log models.LogModel
	global.DB.Take(&log)
	byteData, _ := json.Marshal(log)
	fmt.Println(string(byteData))
}

func main() {
	global.Connect()
	//global.Migrate()
	test()
}
```
```go
package models

import "encoding/json"

type Level int8

const (
	InfoLevel    Level = 1
	WarningLevel Level = 2
	ErrorLevel   Level = 3
)

func (s Level) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case InfoLevel:
		str = "info"
	case WarningLevel:
		str = "warning"
	case ErrorLevel:
		str = "error"
	}
	//return json.Marshal(str)
	return json.Marshal(map[string]any{
		"value": int8(s),
		"label": str,
	})
}

type LogModel struct {
	ID    uint   `json:"id"`
	Title string `gorm:"size:32"`
	Level Level  `json:"level"`
}
```

### 事务
```go
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
```