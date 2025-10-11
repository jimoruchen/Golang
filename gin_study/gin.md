# Golang

## gin

### go内置http库
```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.URL.String())
	if request.Method != "GET" {
		byteData, _ := io.ReadAll(request.Body)
		fmt.Println(string(byteData))
	}
	fmt.Println(request.Header)

	writer.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/index", Index)
	fmt.Println("http server running 127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
```

### 初始gin框架
```go
package main

import "github.com/gin-gonic/gin"

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Index2(c *gin.Context) {
	c.JSON(200, response{
		Code: 0,
		Msg:  "成功",
		Data: map[string]any{},
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	//1.初始化
	r := gin.Default()
	//2.挂载路由
	r.GET("/index", Index2)
	//3.绑定端口，运行
	r.Run(":8080")
}
```

### 响应

#### 响应JSON
```go
package main

import (
	"Golang/gin_study/response/res"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"code": 0,
		//	"msg":  "成功",
		//	"data": gin.H{},
		//})
		res.OkWithMsg(c, "登录成功")
	})

	r.GET("users", func(c *gin.Context) {
		res.OkWithData(c, map[string]any{
			"name": "zhangsan",
		})
	})

	r.POST("users", func(c *gin.Context) {
		res.FailWithMsg(c, "参数错误")
	})

	r.Run(":8080")
}
```
```go
package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var codeMsg = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

func response(c *gin.Context, code int, data any, msg string) {
	c.JSON(200, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context, data any, msg string) {
	response(c, 0, data, msg)
}

func OkWithData(c *gin.Context, data any) {
	Ok(c, data, "成功")
}

func OkWithMsg(c *gin.Context, msg string) {
	Ok(c, gin.H{}, msg)
}

func Fail(c *gin.Context, code int, data any, msg string) {
	response(c, code, data, msg)
}

func FailWithMsg(c *gin.Context, msg string) {
	response(c, 1001, nil, msg)
}

func FailWithCode(c *gin.Context, code int) {
	msg, ok := codeMsg[code]
	if !ok {
		msg = "服务错误"
	}
	response(c, code, nil, msg)
}
```

#### 响应HTML
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("gin_study/response/static/*")  	//加载所有
	r.LoadHTMLFiles("gin_study/response/static/index.html") //加载单个
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", map[string]any{
			"title": "即墨如尘",
		})
	})
	r.Run(":8080")
}
```
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.title}}</title>
</head>
<body>
zhangsan
</body>
</html>
```

#### 文件响应
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename= 5.go")
		c.File("5.go")
	})
	r.Run(":8080")
}
```

#### 静态文件
```html
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("st", "static") //前面是别名，后面是实际路径
	r.StaticFile("abc", "static/download.html")
	r.Run(":8080")
}
```

### 请求

#### 查询参数 Query
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "25")
		keyList := c.QueryArray("key")
		fmt.Println(name, age, keyList)
	})
	r.Run(":8080")
}
```

#### 动态参数 Param
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})
	r.Run(":8080")
}
```

#### 表单参数
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		age, ok := c.GetPostForm("age") //判断有没有传，传的为空OK为true，没传为false
		fmt.Println(name)
		fmt.Println(age, ok)
	})
	r.Run(":8080")
}
```

#### 文件上传
```go
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(fileHeader.Filename)
		fmt.Println(fileHeader.Size)

		file, _ := fileHeader.Open()
		byteData, _ := io.ReadAll(file)
		err = os.WriteFile("xxx.jpg", byteData, 0666)
		fmt.Println(err)
	})
	r.Run(":8080")
}
```
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(fileHeader.Filename)
		fmt.Println(fileHeader.Size)

		//file, _ := fileHeader.Open()
		//byteData, _ := io.ReadAll(file)
		//err = os.WriteFile("xxx.jpg", byteData, 0666)
		//fmt.Println(err)

		err = c.SaveUploadedFile(fileHeader, "./upload/"+fileHeader.Filename)
		fmt.Println(err)
	})
	r.Run(":8080")
}
```

多文件上传
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, fileHeaders := range form.File {
			for _, fileHeader := range fileHeaders {
				c.SaveUploadedFile(fileHeader, "upload/"+fileHeader.Filename)
			}
		}
	})
	r.Run(":8080")
}
```

#### 原始参数 GetRawData
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _raw(c *gin.Context, obj any) (err error) {
	byteData, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(byteData, &obj)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(obj)
	}
	return nil
}

func updateUser(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var user User
	err := _raw(c, &user)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, Response{0, user, "成功"})
}

func main() {
	r := gin.Default()
	r.POST("/", updateUser)
	r.Run(":8080")
}
```

### bind绑定器

#### 查询参数
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBindQuery(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
```

#### 路径参数
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("users/:id/:name", func(c *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			Id   int    `uri:"id"`
		}
		var user User
		err := c.ShouldBindUri(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
```

#### 绑定表单
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("form", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Id   int    `form:"id"`
		}
		var user User
		err := c.ShouldBind(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
```

#### 绑定JSON
```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		fmt.Println(user, err)
	})
	r.Run(":8080")
}
```

#### 参数校验

##### required：必填字段，如：binding:"required", min 最小长度，如：binding:"min=5", max 最大长度，如：binding:"max=10"
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required,min=2,max=6"`
			Age  int    `json:"age"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
```

##### eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`, nefield 不等于其他字段的值
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type User struct {
			Pwd   string `json:"pwd" binding:"required"`
			RePwd string `json:"rePwd" binding:"eqfield=Pwd"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
```

##### 枚举  只能是red 或green, oneof=red green
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"oneof=xxx vvv"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
```

##### 字符串
contains=fengfeng  // 包含fengfeng的字符串
excludes // 不包含
startswith  // 字符串前缀
endswith  // 字符串后缀
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type User struct {
			FileName string `json:"filename" binding:"endswith=.jpg"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
```

##### IP校验
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type User struct {
			IP string `json:"ip" binding:"required,ip"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
```

##### IP数组校验
{
"ipList": ["127.1.2.1","123.1.2.2"]
}
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type User struct {
			IPList []string `json:"ipList" binding:"dive,ip"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)
	})
	r.Run(":8080")
}
```

##### 错误信息显示中文
```go
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
```
```go
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
```

##### 自定义验证器
```go

```

