package pkg

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {

	authenticate, err := BankAuthenticate("AppCode",
		"张三",
		"6225756663322156",
		"34042158962596321")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", authenticate)
	fmt.Printf("%s", authenticate.RespCode)
	fmt.Printf("%s", authenticate.RespCode.ErrorValue())

}
