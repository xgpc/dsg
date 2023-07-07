package dsg

import (
	"github.com/xgpc/dsg/v2/pkg/aes/ecb_aes"
	"github.com/xgpc/dsg/v2/pkg/util"
)

// aes
var _aesKey []byte

func OptionAes(aesKey string) option {
	return func() error {
		_aesKey = []byte(aesKey)
		return nil
	}

}

func key() []byte {
	if len(_aesKey) == 0 {
		panic("keyByte is nil")
	}
	return _aesKey
}

func AESEnCode(data string) []byte {
	return ecb_aes.AESEncrypt([]byte(data), key())
}

func AESDeCode(data []byte) string {
	return string(ecb_aes.AESDecrypt(data, key()))
}

func EnMobile(mobile string) string {
	return mobile[0:3] + "****" + mobile[7:]
}

func CheckMobile(mobile string) bool {
	return util.ValidatePhone(mobile)
}

func EnIDCard(idCard string) string {
	return idCard[0:6] + "********" + idCard[14:]
}

func CheckIDCard(idCard string) bool {
	return util.ValidateIDCard(idCard)
}
