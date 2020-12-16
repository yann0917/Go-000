package router

import (
	. "homework/internal/controllers"

	"github.com/wlxpkg/zwyd/middleware"
	// "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// pprof.Register(r, "debug/pprof")   gin.Logger()
	r.Use(gin.Recovery(), middleware.Recover())
	gin.SetMode(config.Getenv("GIN_MODE", gin.ReleaseMode))

	admin := r.Group("/admin")
	admin.Use(middleware.Casbin())
	{
		ad := admin.Group("/ad")
		{
			ad.POST("/", AddAd)
			ad.GET("/list/:type", AdList)
			ad.PUT("/:id", EditAd)
			ad.GET("/info/:id", AdInfo)
			ad.DELETE("/:id", DelAd)
		}

		app := admin.Group("/app")
		{
			app.GET("/list/:type", GetAppList)
			app.POST("", AddCommonApp)
			app.PUT("/:id", UpdateCommonApp)
		}
	}

	r.GET("ip2address", Ip2Address) // 获取配置

	return r
}
