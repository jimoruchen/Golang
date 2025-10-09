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
