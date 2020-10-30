package main

import (
	"fmt"
	"net/http"
	"redis-haha/router"
	"redis-haha/util"
)

func main() {

	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", util.HTTPPort),
		Handler:        r,
		ReadTimeout:    util.ReadTimeout,
		WriteTimeout:   util.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()

	//conn := Connection{
	//	Type:     "redis",
	//	Host:     "localhost",
	//	port:     "6379",
	//	UserName: "yanghaha",
	//	PassWord: "yanghaha",
	//}
	//
	//str, _ := json.Marshal(conn)
	//fmt.Println(base64.StdEncoding.EncodeToString(str))
	//
	//var tempConn Connection
	//_ = json.Unmarshal(str, &tempConn)
	//fmt.Println(tempConn)

}
