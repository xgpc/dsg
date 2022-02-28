package exce

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"
)

func Write(data string) {
	day := time.Now().Format("2006-01-02")
	dir := "log"
	path := dir + "/" + day + ".log"
	create(dir)

	// 追写文件
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}

	//
	t := time.Now().Format("2006-01-02 15:4:05")

	_, err = f.WriteString(t + " " + data)
	if err != nil {
		log.Println(err)
	}
	_, err = f.Write(debug.Stack())
	if err != nil {
		log.Println(err)
	}
}

func create(path string) {
	if !isExist(path) {
		err := os.MkdirAll(path, 0644)
		if err != nil {
			log.Println(err)
		}
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
