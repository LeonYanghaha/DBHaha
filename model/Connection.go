package model

import (
	"DBHaha/util"
)

// TODO 增加一个添加时间
type Connection struct {
	Id       string `form:"id" json:"Id" xml:"id"`
	NickName string `form:"nickname" json:"Nickname" xml:"nickname"  binding:"required"`
	Type     string `form:"type" json:"Type" xml:"type"  binding:"required"`
	Host     string `form:"host" json:"Host" xml:"host"  binding:"required"`
	Port     int    `form:"port" json:"Port" xml:"port"  binding:"required"`
	UserName string `form:"username" json:"Username" xml:"username" `
	PassWord string `form:"password" json:"Password" xml:"password" `
}

// 加密各个字段
func (conn *Connection) EncryptField() *Connection {

	conn.Type = util.AesEncrypt(conn.Type)
	conn.Host = util.AesEncrypt(conn.Host)
	//conn.Port = util.AesEncrypt(conn.Port)
	conn.UserName = util.AesEncrypt(conn.UserName)
	conn.PassWord = util.AesEncrypt(conn.PassWord)
	return conn
}

// 解密各个字段
func (conn *Connection) DecryptField() *Connection {

	conn.Type = util.AesDecrypt(conn.Type)
	conn.Host = util.AesDecrypt(conn.Host)
	//conn.Port = util.AesDecrypt(strconv.Itoa(conn.Port))
	conn.UserName = util.AesDecrypt(conn.UserName)
	conn.PassWord = util.AesDecrypt(conn.PassWord)
	return conn
}
