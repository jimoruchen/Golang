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

