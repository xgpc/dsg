// Package ecb_aes
// @Author:        asus
// @Description:   $
// @File:          aes_decrypt_test.go
// @Data:          2022/8/1917:12
//
package ecb_aes

import (
	"fmt"
	"github.com/xgpc/dsg"
	"github.com/xgpc/dsg/frame"
	"insurance/models"
	"testing"
)

func TestDemo(t *testing.T) {
	dsg.New("../../config.yaml")
	key := []byte("12345")

	// 调试
	var info models.Demo
	db := frame.MySqlDefault()

	db.Model(&info).Limit(1).First(&info)

	fmt.Println(info)

	info.IDCardData = AESEncrypt([]byte(info.IDCard), key)
	info.MobileData = AESEncrypt([]byte(info.Mobile), key)

	err := db.Save(&info).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}
