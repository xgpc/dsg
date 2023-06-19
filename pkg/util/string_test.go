package util

import (
	"fmt"
	"testing"
)

func TestCamelCaseString(t *testing.T) {
	str := `abc_def_ghi`
	fmt.Println(CamelCaseString(str))
}
