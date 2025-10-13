package main

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type UserInfo struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名必须是 xxx"`
	Age  int    `json:"age" binding:"gte=1,lte=150" msg:"年龄必须在 1-150 之间"`
}

func GetValidMsg(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fieldErr := range ve {
			// 获取结构体字段
			f, exist := reflect.TypeOf(UserInfo{}).FieldByName(fieldErr.Field())
			if exist {
				if msg := f.Tag.Get("msg"); msg != "" {
					return msg
				}
			}
		}
	}
	return err.Error()
}

func GetValidMsg1(err error, obj any) string {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return err.Error()
	}
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for _, fieldErr := range ve {
		f, exist := t.FieldByName(fieldErr.Field())
		if exist {
			if tag := f.Tag.Get("msg"); tag != "" {
				return tag
			}
		}
	}
	return err.Error()
}

func signValid(fl validator.FieldLevel) bool {
	field := fl.Field() //返回值是 reflect.Value 类型
	if field.Kind() != reflect.String {
		return false
	}
	return field.String() == "xxx"
}

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var user UserInfo
		if err := c.ShouldBindJSON(&user); err != nil {
			msg := GetValidMsg1(err, &user)
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  msg,
				"data": nil,
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": user,
		})
	})

	router.Run(":8080")
}
