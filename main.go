package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ClassModel struct {
	Name string `jimoruchen-orm:"name"`
	Id   int    `jimoruchen-orm:"id"`
}

func Find(obj any, query ...any) (sql string, err error) {
	//判断obj必须是结构体
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		err = errors.New("obj must be a struct")
		return
	}

	var where string
	//验证问号个数是否是对等的
	if len(query) > 0 {
		q := query[0]
		qs, ok := q.(string)
		if !ok {
			err = errors.New("query must be string")
			return
		}
		if strings.Count(qs, "?")+1 != len(query) {
			err = errors.New("查询参数不匹配")
		}
		//拼接where
		for _, a := range query[1:] {
			switch s := a.(type) {
			case string:
				qs = strings.Replace(qs, "?", fmt.Sprintf("'%s'", s), 1)
			case int:
				qs = strings.Replace(qs, "?", fmt.Sprintf("%d", s), 1)
			}
		}
		where = "where " + qs
	}
	//拼接所有的有jimoruchen-orm的字段
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ormTag := field.Tag.Get("jimoruchen-orm")
		if ormTag == "" {
			continue
		}
		columns = append(columns, ormTag)
	}

	//算表名，小写机构体的名字加下划线s
	name := strings.ToLower(t.Name()) + "s"
	sql = fmt.Sprintf("select %s from %s %s", strings.Join(columns, ", "), name, where)
	return
}

func main() {
	sql, err := Find(ClassModel{}, "name = ?", "三年一班")
	fmt.Println(sql, err)
	sql, err = Find(ClassModel{}, "id = ? and name = ?", 1, "三年一班")
	fmt.Println(sql, err)
	sql, err = Find(ClassModel{})
	fmt.Println(sql, err)
}
