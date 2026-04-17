package hlog

import (
	"context"

	"github.com/gogf/gf/v2/os/glog"
)

var (
	ctx  = context.TODO()
	path = "hlogs"
)

type LOG_LEVEL struct {
	Level string
}

func (l LOG_LEVEL) Log(i ...interface{}) {

	glog.SetPath(path)
	glog.SetStdoutPrint(false)

	template := glog.File("error-{Ymd}.log")

	switch l.Level {

	case "ERRO":
		template.Error(ctx, i)

	case "INFO":
		template.Info(ctx, i)

	case "DEBU":
		template.Debug(ctx, i)

	default:
		template.Error(ctx, i)
	}

}
