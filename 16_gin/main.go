package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/pingg", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "啦啦啦",
		})
	})
	engine.Run(":1234")
}