package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
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

// 解析文件内容--通过ID获取指定行
func ParseFileGetID(id string) string {
	fileContextList := ParseFile(GetUserInfoFile())
	var lineStr string
	for _, v := range fileContextList {
		if find := strings.Contains(v, id); find {
			lineStr = v
			break
		}
	}
	return lineStr
}

//删除指定的行
func DeleteFileLineById(id string) error {
	fileContextList := ParseFile(GetUserInfoFile())
	removeInfoFile() // 先读取，后删除
	var fileData string
	for _, v := range fileContextList {
		v += "\n"
		if find := strings.Contains(v, id); find {
			continue
		}
		fileData += v
	}
	WriteFile(GetUserInfoFile(), []byte(fileData))
	return nil
}

//修改用户的文件
func UpdateFileLineById(id, contextText string) error {
	fileContextList := ParseFile(GetUserInfoFile())
	removeInfoFile() // 先读取，后删除
	var fileData string
	for _, v := range fileContextList {
		v += "\n"
		if find := strings.Contains(v, id); find {
			v = contextText + "\n"
		}
		fileData += v
	}
	WriteFile(GetUserInfoFile(), []byte(fileData))
	return nil
}

// 解析文件内容---获取所有行
func ParseFile(path string) []string {

	lineContext := ReadFile(path)
	var lineConTextList []string
	for i, v := range lineContext {
		if i == 0 { // 过滤掉文件第一行
			continue
		}
		lineConTextList = append(lineConTextList, v)
	}
	return lineConTextList
}

func removeInfoFile() {
	_ = os.Remove(GetUserInfoFile())
}

// 写入文件
func WriteFile(fileName string, strTest []byte) {

	initFile(fileName)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	fmt.Println("string(strTest))", string(strTest))
	buf := fmt.Sprintf("\n" + string(strTest))
	if _, err := f.WriteString(buf); err != nil {
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

	if _, err := f.WriteString("---DBHaha configuration file, please do not modify it at will---"); err != nil {
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
	fileName := ConfInfo["ConfigFile"]
	home, _ := Home()
	return path.Join(home, fileName)
}

func Home() (string, error) {
	current, err := user.Current()
	if nil == err {
		return current.HomeDir, nil
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
	gateau := os.Getenv("HOMEPATH")
	home := drive + gateau
	if drive == "" || gateau == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	return home, nil
}
