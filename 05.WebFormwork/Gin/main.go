package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	gin.SetMode("debug")

	r := gin.New()
	// 加载中间件，有顺序要求
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 新建一个ping
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。
	r.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// callback 是 x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// 绑定
	r.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	// Query 和 post form
	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// 上传文件
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
		}
		// 上传文件至指定目录
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// 启动并监听在0.0.0.0:3000
	r.Run(":3000")
}
