package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

type config struct {
	DateDue               string `json:"date_due"`
	Oss                   bool   `json:"oss"`
	ExternalControl       bool   `json:"external_control"`
	FormalEnvironmentUrl  string `json:"formal_environment_url"`
	TestingEnvironmentUrl string `json:"testing_environment_url"`
	GroupId               string `json:"group_id"`
}

// PKCS7Padding 使用PKCS7进行填充，IOS也是7
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

// AesCBCEncrypt aes加密，填充模式由key决定，16位，24,32分别对应AES-128, AES-192, or AES-256.源码好像是写死16了
func AesCBCEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()

	rawData = PKCS7Padding(rawData, blockSize)
	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, blockSize+len(rawData))
	//block大小 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func AesCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	//解填充
	encryptData = PKCS7UnPadding(encryptData)
	return encryptData, nil
}

func Encrypt(rawData, key []byte) (string, error) {
	data, err := AesCBCEncrypt(rawData, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Decrypt(rawData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := AesCBCDecrypt(data, key)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

func main() {
	fmt.Println()
	//比如我想把一个json的结构体加密
	str := config{
		"1636559999",
		true,
		false,
		"https://mesp-akk.pdf-in.com/v1/",
		"https://panu-sebv-prod.ife.com/v2/",
		"b70c0a89-a0f7-453e-8e1b-48f2aa83422b",
	}
	//要转化成string格式，通过序列化的方式
	marshal, err := json.Marshal(str)
	if err != nil {
		return
	}
	key := "0123456789012345"
	encrypt, err := Encrypt(marshal, []byte(key))
	if err != nil {
		return
	}
	fmt.Println("密文：\n", encrypt)
	/*
		密文：
		vCJjxp/mMvjCku/RAq+zaq07Pm0uCpAnkBp0Gy5Xt0oPDvYOXROsVfIrx4GVrOz2S0gMuIFwR9QTXFhKC+
		17I8JaBHL27ILOCUcBSZJ4sscvFbffrr8WBfyyDJR8ZuiBvNIcQVteBNm99B1kvNb9H/8WQLBKut1RgCRUV
		LRsNnovQkD0HgOutk2W7i+zAylEzOvHxVGxb//vHwhLovXT1ZEZcaRFsSxNKfKac50tunjxVbTbQStRH8eL
		n6R3RCR07+o36H2VT0cU2tCOQxkb/1AOcfjtNL5geJQ1vmNUNhkaCmD/jOEO3v3IGUn/P0VR1MCUEAoR65BfXrkXoNbDeQ==
	*/

	fmt.Println()

	decrypt, err := Decrypt(encrypt, []byte(key))
	if err != nil {
		return
	}
	fmt.Println("明文：\n", decrypt)
	/*
		明文：
		{"date_due":"1636559999","oss":true,"external_control":false,
		"formal_environment_url":"https://mesp-akk.pdf-in.com/v1/",
		"testing_environment_url":"https://panu-sebv-prod.ife.com/v2/",
		"group_id":"b70c0a89-a0f7-453e-8e1b-48f2aa83422b"}
	*/
}
