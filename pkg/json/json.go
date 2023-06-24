package json

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
)

// replace encoding/json with Iter
func Iter() jsoniter.API {
	return jsoniter.ConfigCompatibleWithStandardLibrary
}

// Decode unmarshal
func Decode(s []byte, data interface{}) (err error) {
	if len(s) < 1 {
		return errors.New("nil")
	}
	err = Iter().Unmarshal(s, &data)
	if err != nil {
		return
	}
	return
}

// Encode marshall
func Encode(data interface{}) (marshal []byte, err error) {
	marshal, err = Iter().Marshal(&data)
	if err != nil {
		return
	}
	return
}
