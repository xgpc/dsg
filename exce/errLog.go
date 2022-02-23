package exce

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"runtime"
	"time"
)

func ErrDeal(err error, args ...interface{}) error {
	num := len(args)
	switch num {
	case 0:
		defer func() {
			if err := recover(); err != nil {
				log.Println("日志写入失败：" + err.(error).Error())
			}
		}()
		pc, _, line, ok := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		if !ok {

		}
		if err != nil {
			t := time.Now().Format("2006-01-02 15:4:05")
			errMsg := fmt.Sprintf("%s at %s:%d Cause by: %s\n", t, f.Name(), line, err.Error())
			write(errMsg)
		}
		return status.Error(CodeSysBusy, "")
	case 1:
		code := args[0]
		return status.Error(codes.Code(code.(int)), "")
	case 2:
		code := args[0]
		msg := args[1]
		return status.Error(codes.Code(code.(int)), msg.(string))
	}
	return nil
}

func write(data string) {
	day := time.Now().Format("2006-01-02")
	dir := "log"
	path := dir + "/" + day + ".log"
	create(dir)

	// 追写文件
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	_, err = f.WriteString(data)
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
