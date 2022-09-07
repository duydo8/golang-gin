package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("static_file", "./image")
	r.Use(AMiddleWare)
	r.GET("/ping", getMethod)
	r.POST("/ping", postMethod)
	r.GET("/detail/:id", getParam)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(context *gin.Context) {
		// Single file
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		context.SaveUploadedFile(file, "./image/uploads/"+file.Filename)

		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	api := r.Group("/api")
	{
		v1 := api.Group("v1")
		v1.Use(myGroupV1MiddleWare())
		{
			v1.GET("a", func(context *gin.Context) {
				context.JSON(200, gin.H{
					"message": "chaocaunhe",
				})
			})
		}
		v2 := api.Group("v2")
		{
			v2.GET("a", func(context *gin.Context) {
				context.JSON(200, gin.H{
					"message": "chaocaunhe",
				})
			})
		}

	}
	//upload

	r.Run(":8080")
}

func getParam(context *gin.Context) {
	name := context.DefaultQuery("name", "duy")

	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}
func postMethod(context *gin.Context) {
	address := context.DefaultPostForm("add", "vietnam")
	context.JSON(http.StatusOK, gin.H{
		"message": "hello " + address + " post message",
		"address": address,
	})
}
func getMethod(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello get message",
	})
}
func AMiddleWare(context *gin.Context) {
	fmt.Println("This is a middleware")
	context.Next()
}

func myGroupV1MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("This is v1 middleware")
		context.Next()
	}
}
