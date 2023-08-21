package dsg

import (
	"testing"
	"time"
)

func TestCreateJwt(t *testing.T) {
	err := OptionJwt("test")()
	if err != nil {
		panic(err)
	}

	toekn := CreateToken(1, time.Hour*24)

	res := ParseToken(toekn)
	if res == nil {
		panic(nil)
	}

	if res.UserID != 1 {
		t.Error("user !=")
	}

	err = res.Validate()
	if err != nil {
		t.Error(err)
	}

}
