package router

import (
	"DBHaha/model"
	"DBHaha/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 删除conn
func RemoveConn(c *gin.Context) {

	cid := c.PostForm("cid")

	checkId := util.ParseFileGetID(cid)
	if checkId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id 不合法"})
		return
	}
	err := util.DeleteFileLineById(cid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "success"})

}

// 更新conn
func UpdateConn(c *gin.Context) {
	var conn model.Connection
	if err := c.ShouldBind(&conn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkId := util.ParseFileGetID(conn.Id)
	if checkId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id 不合法"})
		return
	}
	connByteSlice, ee := json.Marshal(*conn.EncryptField())
	if ee != nil {
		fmt.Println(ee)
		return
	}
	_ = util.UpdateFileLineById(conn.Id, string(connByteSlice))

	c.JSON(http.StatusBadRequest, gin.H{"error": "success"})
	return
}

// 新增　conn
func AddConn(c *gin.Context) {

	var conn model.Connection
	if err := c.ShouldBind(&conn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn.Id = util.GetRandomString(24)
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

// 首页
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
		err := json.Unmarshal([]byte(v), &tempConn)
		if err == nil {
			connList = append(connList, *tempConn.DecryptField())
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"isFirst":  !isFirst,
		"connList": connList,
		"title":    util.ServeName,
	})
}
