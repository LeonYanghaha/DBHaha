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

	// 用户操作相关的router
	r.POST("/user/addConn", AddConn)
	r.POST("/user/RemoveConn", RemoveConn)
	r.POST("/user/UpdateConn", UpdateConn)

	r.GET("/", Index)
	return r
}
