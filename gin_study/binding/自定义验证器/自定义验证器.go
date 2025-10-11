package main

import (
	"errors"
	"fmt"
	"net"
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

var trans ut.Translator

// 初始化验证器与翻译器
func init() {
	// 创建中文翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 获取 Gin 默认的 validator 引擎
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("无法获取 validator 引擎")
	}

	// 注册字段名显示：优先使用 `label` tag，其次 `json` tag
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
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

	// 注册默认中文翻译（required, email 等）
	if err := zh_translations.RegisterDefaultTranslations(v, trans); err != nil {
		panic("注册默认翻译失败: " + err.Error())
	}

	// 注册自定义验证：fip (格式化 IP 校验)
	err := v.RegisterValidation("fip", validateFIP)
	if err != nil {
		return
	}

	// 注册 fip 的中文翻译
	if err := v.RegisterTranslation(
		"fip",
		trans,
		func(ut ut.Translator) error {
			// 添加翻译模板
			return ut.Add("fip", "{0}不是一个有效的IP地址", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			// 实际翻译：{0} 被替换为字段名（来自 RegisterTagNameFunc）
			t, _ := ut.T("fip", fe.Field())
			return t
		},
	); err != nil {
		panic("注册 fip 翻译失败: " + err.Error())
	}
}

// 自定义 IP 校验函数
func validateFIP(fl validator.FieldLevel) bool {
	// 获取字段值
	value := fl.Field().Interface()
	if str, ok := value.(string); ok && str != "" {
		// 非空字符串才校验
		return net.ParseIP(str) != nil
	}
	// 空值视为通过（如果需要非空，应配合 required）
	return true
}

// ValidateErr 将验证错误转换为 map[string]string
func ValidateErr(err error) map[string]string {
	var errors1 = make(map[string]string)

	// 判断是否为 validator.ValidationErrors
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		// 不是验证错误，返回通用错误
		errors1["error"] = err.Error()
		return errors1
	}

	// 遍历每个字段错误
	for _, e := range validationErrors {
		field := e.Field() // 来自 RegisterTagNameFunc 的名字（label/json）
		message := e.Translate(trans)
		errors1[field] = message
	}

	return errors1
}

// User 用户结构体示例
type User struct {
	Name  string `json:"name" binding:"required" label:"用户名"`
	Email string `json:"email" binding:"required,email" label:"邮箱地址"`
	Ip    string `json:"ip" binding:"fip" label:"IP地址"`
}

func main() {
	// 设置 Gin 为 Release 模式（可选）
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// POST /user 示例
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			// 参数验证失败
			c.JSON(http.StatusOK, gin.H{
				"code": 7,
				"msg":  "验证错误",
				"data": ValidateErr(err),
			})
			return
		}

		// 验证通过
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": user,
		})
	})

	// 启动服务器
	fmt.Println("服务器已启动：http://localhost:8080")
	r.Run(":8080")
}
