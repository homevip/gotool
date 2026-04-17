package hlog

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	// // 调试信息
	// var (
	// 	ctx  = context.TODO()
	// 	path = "./glog" // 日志存储路径
	// )

	// glog.SetPath(path)
	// glog.SetStdoutPrint(false)

	// glog.File("error-{Ymd}.log").Error(ctx, "文件名称支持带gtime日期格式")

	var (
		l = LOG_LEVEL{}
	)
	err := "error ..."
	errInfo := fmt.Sprintf("异常信息:%v", err)
	l.Log(errInfo)

}
