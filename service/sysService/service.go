package sysService

import (
	"time"
)

var SysVersion int64

func GetSetSysVersion() int64 {
	if SysVersion == 0 {
		SysVersion = time.Now().Unix()
	}
	return SysVersion
}

func InitSysVersion() {
	SysVersion = time.Now().Unix()
}
