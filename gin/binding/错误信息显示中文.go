package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var translator ut.Translator

func init() {
	// 初始化中文翻译器
	zhLocale := zh.New()
	uni := ut.New(zhLocale, zhLocale)
	translator, _ = uni.GetTranslator("zh")

	// 获取 Gin 默认的 validator 引擎
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("failed to get validator from gin")
	}

	// 注册 tagName 函数：优先使用 label tag
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label != "" {
			return label
		}
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			return strings.Split(jsonTag, ",")[0]
		}
		return field.Name
	})

	// 将中文翻译注册到 Gin 使用的 validator 实例上
	if err := zh_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		panic("failed to register translations: " + err.Error())
	}
}

// TranslateValidationError 统一返回验证错误信息
func TranslateValidationError(err error) string {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var messages []string
		for _, e := range errs {
			messages = append(messages, e.Translate(translator))
		}
		return strings.Join(messages, "; ")
	}
	return err.Error()
}

type User struct {
	Name  string `json:"name" binding:"required" label:"用户名"`
	Email string `json:"email" binding:"required,email" label:"邮箱地址"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(http.StatusOK, TranslateValidationError(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s! Your email is %s.", user.Name, user.Email),
		})
	})

	r.Run()
}
