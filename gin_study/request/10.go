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
