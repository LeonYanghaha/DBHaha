package router

import (
	"DBHaha/model"
	"DBHaha/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddConn(c *gin.Context) {

	var conn model.Connection
	if err := c.ShouldBind(&conn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn.Id = util.GetRandomString(10)
	fmt.Println("conn", conn)
	connByteSlice, ee := json.Marshal(*conn.EncryptField())
	if ee != nil {
		fmt.Println(ee)
	}
	fmt.Println("connByteStr", string(connByteSlice))

	util.WriteFile(util.GetUserInfoFile(), connByteSlice)
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
	var connList []model.Connection
	for _, v := range fileContext {
		var tempConn model.Connection
		err := json.Unmarshal([]byte(v), tempConn)
		if err == nil {
			connList = append(connList, *tempConn.DecryptField())
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"isFirst":  !isFirst,
		"connList": fileContext,
		"title":    util.ServeName,
	})
}
