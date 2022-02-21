package cryptService

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

func GenKey(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	key := base64.StdEncoding.EncodeToString(b)
	if len(key) < 16 {
		key += strings.Repeat("0", 16-len(key))
	} else if len(key) > 16 {
		key = key[:16]
	}

	fmt.Println(key)
	decodeString, err := base64.StdEncoding.DecodeString(key)
	fmt.Println(len(decodeString))

	return key, nil
}

func GenKeyByte(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
