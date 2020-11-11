package util

import (
	"math/rand"
	"time"
)

func GetResData(msg string, code int, data interface{}) map[string]interface{} {
	resMap := make(map[string]interface{})
	resMap["msg"] = msg
	resMap["code"] = code
	resMap["data"] = data
	return resMap
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
