package main

import (
	"dsg/service/cryptService"
	"fmt"
)

func main() {
	//加密
	data := []byte("hello world")
	publicKey := cryptService.ReadRSAKey("public.pem")
	privateKey := cryptService.ReadRSAKey("private.pem")
	encrypt := cryptService.RSAEncrypt(data, publicKey)

	// 解密
	decrypt := cryptService.RSADecrypt(encrypt, privateKey)
	fmt.Println("decrypt:", string(decrypt))
}
