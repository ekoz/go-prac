package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    nil,
	})
}

type User struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

// 权限过滤中间件
func securityFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("security filter in..")
		c.Next()
		fmt.Println("security filter out..")
	}
}

func main() {
	r := gin.Default()

	// 绑定中间件
	r.Use(securityFilter())

	r.GET("/hello", sayHello)

	// 获取路径中的参数
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": c.Query("name"),
			"age":  c.Query("age"),
		})
	})

	// 获取 uri 中的参数
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"year":  c.Param("year"),
			"month": c.Param("month"),
		})
	})

	// 获绑定结构体
	r.GET("/login", func(c *gin.Context) {
		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求参数错误 " + err.Error(),
			})
			return
		}
		fmt.Printf("%v\n", user)
		c.JSON(http.StatusOK, gin.H{
			"message": user.Name + ` 登录成功`,
		})
	})

	// 上传文件
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		fh, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求参数错误 " + err.Error(),
			})
			return
		}

		// dst := fmt.Sprintf("./%s", fh.Filename)
		dst := path.Join("./", fh.Filename)
		c.SaveUploadedFile(fh, dst)

		c.JSON(http.StatusOK, gin.H{
			"message": "文件上传成功",
		})

	})

	// 请求重定向
	r.GET("/baidu", func(c *gin.Context) {

		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 路由组
	bookGroup := r.Group("/book")
	// {} 非必须，推荐写法
	{
		bookGroup.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get book list",
			})
		})

		bookGroup.GET("/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get book detail",
			})
		})

		bookGroup.GET("/page", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "get book page",
			})
		})
	}

	// 未配置的路由统一走这里
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404 not found",
		})
	})

	r.Run(":8080")
}
