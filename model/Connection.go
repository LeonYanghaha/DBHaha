package model

type Connection struct {
	Id       string `form:"id" json:"id" xml:"id"`
	NickName string `form:"nickname" json:"nickname" xml:"nickname"  binding:"required"`
	Type     string `form:"type" json:"type" xml:"type"  binding:"required"`
	Host     string `form:"host" json:"host" xml:"host"  binding:"required"`
	Port     string `form:"port" json:"port" xml:"port"  binding:"required"`
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	PassWord string `form:"password" json:"password" xml:"password"  binding:"required"`
}
