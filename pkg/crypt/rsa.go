package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/xgpc/dsg/exce"
	"os"
)

func ReadRSAKey(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	//block, _ := pem.Decode(buf)
	return buf
}

// RSAEncrypt RSA加密
// plainText 要加密的数据
func RSAEncrypt(plainText []byte, key []byte) []byte {
	////pem解码
	block, _ := pem.Decode(key)
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		exce.ThrowSys(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		exce.ThrowSys(err)
	}
	return cipherText
}

// RSADecrypt RSA解密
// cipherText 需要解密的byte数据
func RSADecrypt(cipherText []byte, key []byte) []byte {
	//pem解码
	decode, _ := pem.Decode(key)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(decode.Bytes)
	if err != nil {
		exce.ThrowSys(err)
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		exce.ThrowSys(err)
	}
	//返回明文
	return plainText
}
