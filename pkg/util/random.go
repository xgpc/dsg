// Package service
// @Author:        asus
// @Description:   $
// @File:          random
// @Data:          2021/12/2318:44
//
package util

import (
	"math/rand"
	"time"
)

var randomChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var randomNumber = []rune("0123456789")

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

//随机生成字符串
func RandStr(length int, letter []rune) string {
	b := make([]rune, length)
	randomCharsLen := len(letter)

	for i := range b {
		b[i] = letter[rander.Intn(randomCharsLen)]
	}
	return string(b)
}

func RandomStr(length int) string {
	return RandStr(length, randomChars)
}

func RandomNumber(length int) string {
	return RandStr(length, randomNumber)
}

func RandomInt(length int) int {
	return rand.Intn(length)
}
