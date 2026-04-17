package hlog

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 全局 Logger
var Logger *zap.Logger

// InitZap 初始化日志
func InitLogger() {
	atomLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	// 文件切割
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/server.log", path),
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	})

	// 控制台输出
	consoleWriter := zapcore.AddSync(os.Stdout)

	// ====================== 关键修复：分开配置编码器 ======================
	// 1. 文件编码器：无颜色
	fileEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 无颜色
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 2. 控制台编码器：有颜色
	consoleEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 文件输出核心
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(fileEncoderConfig),
		fileWriter,
		atomLevel,
	).With([]zap.Field{zap.String("app", "noaddr-management")})

	// 控制台输出核心
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		consoleWriter,
		atomLevel,
	)

	// 合并
	combinedCore := zapcore.NewTee(fileCore, consoleCore)

	// 生成 logger
	Logger = zap.New(
		combinedCore,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	zap.ReplaceGlobals(Logger)
	Logger.Info("日志初始化成功 ✅")
}

// CustomTimeEncoder 自定义时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
