package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	"github.com/yann0917/Go-000/Week02/controller"
)

// InitRouter gin router
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger()) // gin.Logger(),
	gin.SetMode(config.Getenv("GIN_MODE", gin.ReleaseMode))

	// 前台接口
	front := r.Group("/")
	front.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "pong",
		})
	})

	user := front.Group("user")
	{
		user.GET("/:id", controller.GetUser)
	}

	return r
}
