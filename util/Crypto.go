package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
)

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
func AesEncrypt(origDataStr string) string {
	origData := []byte(origDataStr)
	key := []byte(Secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return hex.EncodeToString(crypted)
}

//AES解密
func AesDecrypt(cryptedStr string) string {

	key := []byte(Secret)
	crypted, _ := hex.DecodeString(cryptedStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return string(origData)
}
