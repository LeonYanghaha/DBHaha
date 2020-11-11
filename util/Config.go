package util

var ConfInfo = make(map[string]string)

func init() {
	ConfInfo["Port"] = "8000"
	ConfInfo["ServeName"] = "DBHaha ~"
	//debug or release
	ConfInfo["RunMode"] = "debug"
	// 用户的数据文件
	ConfInfo["ConfigFile"] = "DBHaha/info.data"
	//secret 的长度有长度限制，必须是 16 字节-AES-128  24 字节-AES-192  32 字节-AES-256
	ConfInfo["Secret"] = "0f90023fc9ae101e"
	ConfInfo["ReadTimeout"] = "60"
	ConfInfo["WriteTimeout"] = "60"
}
