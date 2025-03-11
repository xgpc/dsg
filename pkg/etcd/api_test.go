package etcd

import (
	"fmt"
	"strings"
	"testing"
)

func Test_parseServiceKey1(t *testing.T) {
	str := "/service/basic/127.0.0.1:56716"

	strList := strings.Split(str, "/")
	fmt.Println(strList)
}
