package main

import (
	"DBHaha/router"
	"DBHaha/util"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//-------------------------------------------------------------------------------------------
	r := router.InitRouter()

	port, _ := strconv.Atoi(util.ConfInfo["Port"])
	tempReadTimeout, _ := strconv.Atoi(util.ConfInfo["ReadTimeout"])
	ReadTimeout := time.Duration(tempReadTimeout) * time.Second

	tempWriteTimeout, _ := strconv.Atoi(util.ConfInfo["WriteTimeout"])
	WriteTimeout := time.Duration(tempWriteTimeout) * time.Second

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()

	//-------------------------------------------------------------------------------------------
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
