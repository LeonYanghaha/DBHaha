package router

import (
	"DBHaha/util"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.LoadHTMLGlob("templates/page/*")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(util.RunMode)

	// 用户操作相关的router
	r.GET("/", Index)
	r.POST("/user/addConn", AddConn)
	r.POST("/user/RemoveConn", RemoveConn)
	r.POST("/user/UpdateConn", UpdateConn)

	return r
}
