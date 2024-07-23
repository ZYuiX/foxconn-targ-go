package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

/**
使用gin框架完成一个简单的api，
1.页面路由
2.参数传递
3.调用后能正确相应
*/

func main() {
	//启动一个默认路由
	router := gin.Default() // 创建带有默认中间件的 Gin 路由器

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080") // 监听并服务于 0.0.0.0:8080

}