package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redis-haha/model"
	"redis-haha/util"
	"strings"
)

func AddConn(c *gin.Context) {

	var conn model.Connection
	if err := c.ShouldBind(&conn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn.Id = util.GetRandomString(10)
	fmt.Println("conn", conn)
	connByteSlice, ee := json.Marshal(conn)
	if ee != nil {
		fmt.Println(ee)
	}
	connByteStr := strings.Replace(string(connByteSlice), "\"", "'", -1)
	fmt.Println("util.Secret", util.Secret)
	fmt.Println("connByteStr", connByteStr)
	str := string(util.DesEncrypt([]byte(connByteStr), []byte(util.Secret)))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println("WriteFile", str)
	fmt.Println("WriteFile", string(str))
	util.WriteFile(util.GetUserInfoFile(), []byte(str))
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
	fmt.Println("fileContext", fileContext)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"isFirst":  !isFirst,
		"connList": fileContext,
		"title":    util.ServeName,
	})
}
