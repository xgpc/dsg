package aesService

import "github.com/xgpc/dsg/service/aesService/ecb_aes"

var keyByte []byte

func Init(key string) {
	keyByte = []byte(key)
}

func key() []byte {
	if len(keyByte) == 0 {
		panic("keyByte is nil")
	}
	return keyByte
}

func EnCode(data string) []byte {
	return ecb_aes.AESEncrypt([]byte(data), key())
}

func DeCode(data []byte) string {
	return string(ecb_aes.AESDecrypt(data, key()))
}

func EnMobile(mobile string) string {
	return mobile[0:3] + "****" + mobile[7:]
}

func EnIDCard(idCard string) string {
	return idCard[0:6] + "********" + idCard[14:]
}
