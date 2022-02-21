package main

import (
	"encoding/base64"
	"fmt"
	"github.com/xgpc/dsg/service/cryptService"
)

func main() {

	//生成key和iv
	text := "13518757974" // 你要加密的数据
	//
	//AesKey := []byte("1LERnLNmihJrESbE") // 对称秘钥长度必须是16的倍数
	AesKey := "mIEC&BK5geZnr%X#"
	iv := "Fapaecr!JD!53r5s"
	//
	fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
	fmt.Printf("iv: %s 长度: %d\n", iv, len(iv))
	encrypted, err := cryptService.AesEncrypt([]byte(text), []byte(AesKey), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后: %s\n", base64.StdEncoding.EncodeToString(encrypted))
	origin, err := cryptService.AesDecrypt(encrypted, []byte(AesKey), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后明文: %s\n", string(origin))

	////生成key和iv
	//text := "13518757974" // 你要加密的数据
	//
	////AesKey := []byte("1LERnLNmihJrESbE") // 对称秘钥长度必须是16的倍数
	//AesKey := cryptService.GenKeyByte(16)
	//iv := cryptService.GenKeyByte(16)
	//
	//fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
	//fmt.Printf("秘钥: %v\n", AesKey)
	//fmt.Printf("iv: %s 长度: %d\n", iv, len(iv))
	//encrypted, err := cryptService.AesEncrypt([]byte(text), AesKey, iv)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("加密后: %s\n", base64.StdEncoding.EncodeToString(encrypted))
	//decodeString, err2 := base64.StdEncoding.DecodeString("GIQz")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//
	//k, err2 := base64.StdEncoding.DecodeString("M0g0NHVQVXc1dHZ2eVRLbA==")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//i, err2 := base64.StdEncoding.DecodeString("M0g0NHVQVXc1dHZ2eVRLbA==")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//
	//origin, err := cryptService.AesDecrypt(decodeString, k, i)
	//if err != nil {
	//	panic(err)
	//}
	//encodeToString := base64.StdEncoding.EncodeToString(AesKey)
	//toString := base64.RawStdEncoding.EncodeToString(iv)
	//
	//fmt.Println(encodeToString)
	//fmt.Println(toString)
	//fmt.Println(AesKey)
	//fmt.Println(string(AesKey))
}
