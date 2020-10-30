package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redis-haha/model"
	"redis-haha/util"
)

func AddConn(c *gin.Context) {

	var conn model.Connection
	if err := c.ShouldBind(&conn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connByteSlice, _ := json.Marshal(conn)
	secretByteSlice := []byte(util.Secret)
	str, err := util.AesEncrypt(connByteSlice, secretByteSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	util.WriteFile(util.GetUserInfoFile(), str)
	c.JSON(http.StatusBadRequest, gin.H{"error": "success"})
	return
}

func Index(c *gin.Context) {
	filePath := util.GetUserInfoFile()
	var isFirst = false

	isFirst = util.CheckFileExist(filePath)
	// 第一次使用，尚未建立用户文件的情况
	if !isFirst {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"isFirst": !isFirst,
			"title":   util.ServeName,
		})
		return
	}

	fileContext := util.ParseFile(filePath)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"isFirst":  !isFirst,
		"connList": fileContext,
		"title":    util.ServeName,
	})
}
