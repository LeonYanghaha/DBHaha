package util

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"redis-haha/model"
	"runtime"
	"strings"
)

// 读取文件
func ReadFile(path string) []string {
	fi, err := os.Open(path)
	var lineContext []string
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return lineContext
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if string(a) != "" {
			lineContext = append(lineContext, string(a))
		}
	}
	return lineContext
}

// 解析文件内容
func ParseFile(path string) []model.Connection {

	lineContext := ReadFile(path)
	fmt.Println(lineContext)
	var connList []model.Connection

	for i, v := range lineContext {
		var tempConn model.Connection
		if i == 0 {
			continue
		}
		jsonStr, err := AesDecrypt(v, Secret)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println("----------------", hex.EncodeToString(jsonStr))
		err = json.Unmarshal(jsonStr, &tempConn)
		if err != nil {
			connList = append(connList, tempConn)
		}
	}
	return connList
}

// 写入文件
func WriteFile(fileName string, strTest []byte) {

	initFile(fileName)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(hex.EncodeToString(strTest) + "\n"); err != nil {
		log.Println(err)
	}
}

//初始化用户的配置文件
func initFile(fileName string) {

	dir, _ := path.Split(fileName)

	// 如果文件存在，直接返回就好了
	if CheckFileExist(fileName) {
		return
	}
	// 如果文件不存在，就创建文件，并写入默认提示信息
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	_, _ = os.Create(fileName)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString("---DBHaha configuration file, please do not modify it at will---" + "\n"); err != nil {
		log.Println(err)
	}
}

//检查文件是否存在
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 获取当前用户的资料目录
func GetUserInfoFile() string {
	fileName := ConfigFile
	home, _ := Home()
	return home + "/" + fileName
}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}
	if "windows" == runtime.GOOS {
		return homeWindows()
	}
	return homeUnix()
}

func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
