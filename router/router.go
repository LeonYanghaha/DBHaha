package router

import (
	"DBHaha/util"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./templates/static")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(util.ConfInfo["RunMode"])

	r.GET("/", Index)

	// 用户操作相关的router
	userRouter := r.Group("/user")
	{
		userRouter.POST("/addConn", AddConn)
		userRouter.POST("/RemoveConn", RemoveConn)
		userRouter.POST("/UpdateConn", UpdateConn)
	}

	// redis
	redisRouter := r.Group("/redis")
	{
		redisRouter.GET("/open/:cid", OpenRedis)
	}

	return r
}
