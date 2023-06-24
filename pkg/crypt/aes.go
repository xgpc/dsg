package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

// AesEncrypt AES加密CRT
func AesEncrypt(plainText, key []byte, iv []byte) ([]byte, error) {
	if len(key) != 16 || len(iv) != 16 {
		return nil, errors.New("长度必须为16位")
	}
	//1.指定使用的加密aes算法
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(block, iv)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)
	return cipherText, nil
}

// AesDecrypt AES解密CRT
func AesDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	//decodeString, err := base64.StdEncoding.DecodeString(cipherText)
	if len(key) != 16 || len(iv) != 16 {
		return nil, errors.New("长度必须为16位")
	}
	//1.指定算法:aes
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("err:::", err)
		return nil, err
	}
	//2.返回一个计数器模式的、底层采用block生成key流的Stream接口，初始向量iv的长度必须等于block的块尺寸。
	stream := cipher.NewCTR(block, iv)
	//3.解密操作
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)
	return plainText, nil
}
