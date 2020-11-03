package main

import (
	"fmt"
	"os"
	"redis-haha/util"
)

func main() {
	//-------------------------------------------------------------------------------------------
	//r := router.InitRouter()
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", util.HTTPPort),
	//	Handler:        r,
	//	ReadTimeout:    util.ReadTimeout,
	//	WriteTimeout:   util.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//_ = s.ListenAndServe()

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

	//-------------------------------------------------------------------------------------------
	//text :=" {'id':'sccs69oro5','nickname':'localhost-redis','type':'redis','host':'localhost','port':'6379','username':'apollo','password':'1234567809'}"
	//AesKey := []byte("0f90023fc9ae101e")
	//fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
	//encrypted, _ := util.AesEncrypt([]byte(text), AesKey)
	//fmt.Println(hex.EncodeToString(encrypted))
	//
	//str, _ := util.AesDecrypt([]byte(string(encrypted)), AesKey)
	//fmt.Print(string(str))

	//fmt.Println("des 加解密")
	//key := []byte("1234abdd")
	//src := []byte("{'name':'yk','age':'12','address':'bj'}")
	//cipherText := util.DesEncrypt(src, key)
	//fmt.Println(string(cipherText))
	//plainText := util.DesDecrypt([]byte(string(cipherText)), key)
	//fmt.Printf("解密之后的数据: %s\n", string(plainText))
	//
	////fmt.Println("aes 加解密 ctr模式 ... ")
	////key1 := []byte("1234abdd12345678")
	////cipherText = util.AesEncrypt(src, key1)
	////plainText = util.AesDecrypt(cipherText, key1)
	////fmt.Printf("解密之后的数据: %s\n", string(plainText))
	//
	//file, err := os.Open(util.GetUserInfoFile())
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	////fmt.Println(file)
	//payload := cat(file)
	//fmt.Println(string(payload))
	//check(errs)
	//_, errs = fo.Write(payload)

	conn := "{'name':'yk','age':'12','address':'bj'}"
	key := "0f90023fc9ae101e"

	secretSlice,_ := util.AesEncrypt([]byte(conn), []byte(key))
	fmt.Println("secretSlice", secretSlice)
	fmt.Println("secretSlice", string(secretSlice))

	util.WriteFile(util.GetUserInfoFile(), secretSlice)

	file := util.ReadFile(util.GetUserInfoFile())
	fmt.Println(file)
	fmt.Println(util.AesDecrypt([]byte(file[0]),[]byte(key) ))

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cat(f *os.File) []byte {
	var payload []byte
	for {
		buf := make([]byte, 1024)
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return payload
		case nr > 0:
			payload = append(payload, buf...)
		}
	}

}
