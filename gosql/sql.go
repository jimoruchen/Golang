package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:200088@tcp(127.0.0.1:3303)/gorm_db?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败 %s", err)
	}
	fmt.Println(db)
}
