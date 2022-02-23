package exce

import (
	"encoding/json"
	"github.com/xgpc/util"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/xgpc/util/fileU"
)

const (
	LogLvInfo        = "info"
	LogLvError       = "error"
	LogDriverConsole = "console"
	LogDriverLocal   = "local"
	LogDriverServer  = "server"
)

var logConfig = struct {
	Open   bool
	Driver string
	Host   string
	Pwd    string
	Level  string
	Email  string
}{
	false,
	"",
	"",
	"",
    LogLvInfo,
	"",
}

var (
	filePath string
	fileObj  *os.File
)

func LogInfo(msg string, data interface{}) {
	go logHandle(LogLvInfo, msg, &data)
}

func LogError(msg string, data interface{}) {
	go logHandle(LogLvError, msg, &data)
}

func LogException(msg string, data ...interface{}) {
	var debugArr []string
	if len(data) > 0 {
		bt, _ := util.JsonEncode(data[0])

		debugArr = append(debugArr, string(bt))
	}
	debugArr = append(debugArr, string(debug.Stack()))
	var v interface{} = strings.Join(debugArr, "\n")
	logHandle(LogLvError, msg, &v)
}

func logHandle(level string, msg string, data *interface{}) {
	//switch env.LogDriver() {
	//case "":
	//	return
	//case LogDriverConsole:
	//	logConsole(level, msg, data)
	//case LogDriverLocal:
	//	logFile(level, msg, data)
	//case LogDriverServer:
	//	logServer(level, msg, data)
	//}
}

func logConsole(level string, msg string, content *interface{}) {
	log.Printf("%s\r\n%v\r\n", "["+level+"]"+" "+msg, *content)
}

func logFile(level string, msg interface{}, content *interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("日志写入失败：" + err.(error).Error())
		}
	}()

	//go logConsole(level, msg, content)

	day := util.TimeYmdNow()
	target := "./log"
	f := target + "/" + day + ".log"

	if filePath != f {
		filePath = f
		// 关闭文件
		if fileObj != nil {
			_ = fileObj.Close()
		}
		// 建立目录
		if err := fileU.MkDir(target); err != nil {
			panic(err.Error())
		}
		// 创建文件
		if err := fileU.MkFile(filePath); err != nil {
			panic(err.Error())
		}

		// 追写文件
		ff, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		fileObj = ff
	}

	prefix := time.Now().Format("[2006-01-02 15:04:05]") + " [" + level + "] "

	marshal, _ := json.Marshal(msg)

	_, err := fileObj.Write([]byte(prefix + string(marshal) + "\n"))
	if err != nil {
		return
	}

	// 数据转成JSON字符串
	//switch v := (*content).(type) {
	//case string:
	//	buf := []byte(prefix + msg + "\n" + v + "\n\n")
	//	_, _ = fileObj.Write(buf)
	//default:
	//	j, err := json.Marshal(content)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	buf := []byte(prefix + msg + "\n" + string(j) + "\n\n")
	//	_, _ = fileObj.Write(buf)
	//}
}

func logServer(level string, msg string, content *interface{}) {
	// todo:
	logConsole(level, msg, content)
	logConsole(LogLvError, "暂不支持Server方式写日志", nil)
}
