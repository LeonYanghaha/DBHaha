package router

import (
	"DBHaha/model"
	"DBHaha/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
)

func OpRedis(c *gin.Context) {
	cid := c.Param("cid")
	cmd := c.Param("cmd")
	if cid == "" || cmd == "" {
		c.JSON(http.StatusBadRequest, util.GetResData("id 不能为空", nil))
		return
	}

	connString := util.ParseFileGetID(cid)
	var conn model.Connection
	if err := json.Unmarshal([]byte(connString), &conn); err != nil {
		c.JSON(http.StatusInternalServerError, util.GetResData("文件解析错误", nil))
		return
	}
	conn = *conn.DecryptField()
	util.ExecCmd(conn.Id, "", "")
	c.JSON(http.StatusInternalServerError, util.GetResData("success", nil))

}

func OpenRedis(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, util.GetResData("id 不能为空", nil))
		return
	}

	connString := util.ParseFileGetID(cid)
	var conn model.Connection
	if err := json.Unmarshal([]byte(connString), &conn); err != nil {
		c.JSON(http.StatusInternalServerError, util.GetResData("文件解析错误", nil))
		return
	}

	if util.GetCtx(cid) {
		c.HTML(http.StatusOK, "redis.html", gin.H{
			"conn":  conn,
			"title": util.ConfInfo["ServeName"] + conn.NickName,
		})
		return
	}

	conn = *conn.DecryptField()

	redisInfo := util.RedisInfo{
		Addr:      conn.Host + ":" + strconv.Itoa(conn.Port),
		Password:  conn.PassWord,
		DB:        0,
		MyClient:  redis.Client{},
		MyContext: nil,
	}
	redisInfo = util.GetRedisClient(redisInfo)

	util.PutCtx(conn.Id, redisInfo)

	c.HTML(http.StatusOK, "redis.html", gin.H{
		"conn":  conn,
		"title": util.ConfInfo["ServeName"] + conn.NickName,
	})

}
