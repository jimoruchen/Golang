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

