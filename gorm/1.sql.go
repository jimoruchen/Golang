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
