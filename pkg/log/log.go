package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zap log日志
// error logger
var errorLogger *zap.SugaredLogger

type levelType string

const DEBUG levelType = "debug"
const INFO levelType = "info"
const WARN levelType = "warn"
const ERROR levelType = "error"
const DPANIC levelType = "dpanic"
const PANIC levelType = "panic"
const FATAL levelType = "fatal"

var levelMap = map[levelType]zapcore.Level{
	DEBUG:  zapcore.DebugLevel,
	INFO:   zapcore.InfoLevel,
	WARN:   zapcore.WarnLevel,
	ERROR:  zapcore.ErrorLevel,
	DPANIC: zapcore.DPanicLevel,
	PANIC:  zapcore.PanicLevel,
	FATAL:  zapcore.FatalLevel,
}

// init 考虑如果用户没有显示调用Init则默认调用init 防止报错
func init() {
	filePath := "default.log"

	level := getLoggerLevel(ERROR) //日志等级
	hook := lumberjack.Logger{
		Filename: filePath, // 日志文件路径
		MaxSize:  30,       // 每个日志文件保存的最大尺寸 单位：M
		//LocalTime: true,
		MaxBackups: 100,  // 日志文件最多保存多少个备份
		MaxAge:     30,   // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}
	syncWriter := zapcore.AddSync(&hook)

	// 判断日志等级
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level)),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

func Init(configLogLevel levelType, configFilePath string) {
	filePath := "default.log"
	// configFilePath 文件夹位置

	if configFilePath != "" {
		if strings.HasSuffix(configFilePath, "/") { //判断是否以/结尾
			filePath = configFilePath + filePath
		} else {
			filePath = configFilePath + "/" + filePath
		}
	}

	fmt.Printf("log.level:%s\n", configLogLevel)
	level := getLoggerLevel(configLogLevel) //日志等级
	hook := lumberjack.Logger{
		Filename: filePath, // 日志文件路径
		MaxSize:  128,      // 每个日志文件保存的最大尺寸 单位：M
		//LocalTime: true,
		MaxBackups: 100,  // 日志文件最多保存多少个备份
		MaxAge:     30,   // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}
	syncWriter := zapcore.AddSync(&hook)

	// 判断日志等级
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level)),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

func getLoggerLevel(level levelType) zapcore.Level {
	if level, ok := levelMap[level]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
