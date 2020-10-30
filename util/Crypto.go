package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//text := "yanghaha{}{}{}[]()<>,."
//AesKey := []byte("0f90023fc9ae101e")
//fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
//encrypted, _ := util.AesEncrypt([]byte(text), AesKey)
//fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
//
//str, _ := util.AesDecrypt(encrypted, AesKey)
//fmt.Print(string(str))

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES解密
func AesDecrypt(cryptedStr, keyStr string) ([]byte, error) {

	crypted, _ := base64.StdEncoding.DecodeString(cryptedStr)
	key, _ := base64.StdEncoding.DecodeString(keyStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
