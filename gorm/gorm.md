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

#### 