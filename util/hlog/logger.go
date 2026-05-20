package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 包内私有变量,保证全局唯一且外部不可直接修改
var (
	appLogger     *zap.Logger // 全局日志实例
	appLoggerOnce sync.Once   // 保证日志只初始化一次
	appLoggerErr  error       // 初始化错误信息
)

// InitLogger 初始化日志(线程安全,仅执行一次)
// 建议在程序启动时(配置加载完成后)调用
func InitLogger() error {
	appLoggerOnce.Do(func() {
		// 1. 基础配置 - 日志级别
		atomLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

		// 2. 文件切割配置
		today := time.Now().Format("2006-01-02") // 日期格式：2025-09-07
		logPath := fmt.Sprintf("hlogs/server-%s.log", today)
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath, // 每天一个文件
			MaxSize:    128,     // 单个文件最大 128MB
			MaxBackups: 30,      // 最多保留 30 个文件
			MaxAge:     30,      // 最多保留 30 天
			Compress:   true,    // 自动压缩
			LocalTime:  true,    // 使用本地时间切割
		})

		// 3. 控制台输出配置
		consoleWriter := zapcore.AddSync(os.Stdout)

		// 4. 分开配置编码器 (文件:无颜色 / 控制台:有颜色)
		// 4.1 文件编码器配置
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

		// 4.2 控制台编码器配置
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

		// 5. 构建输出核心
		// 5.1 文件输出核心
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(fileEncoderConfig),
			fileWriter,
			atomLevel,
		).With([]zap.Field{zap.String("app", Get().App.Name)})

		// 5.2 控制台输出核心
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(consoleEncoderConfig),
			consoleWriter,
			atomLevel,
		)

		// 6. 合并输出核心 (文件 + 控制台)
		combinedCore := zapcore.NewTee(fileCore, consoleCore)

		// 7. 生成最终的logger实例
		appLogger = zap.New(
			combinedCore,
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		)

		// 8. 替换zap全局logger
		zap.ReplaceGlobals(appLogger)

		fmt.Println("✅ 日志初始化成功")
	})

	return appLoggerErr
}

// GetLogger 获取全局日志实例
// 注意:调用前必须确保InitLogger()已执行且无错误,否则返回nil
func GetLogger() *zap.Logger {
	return appLogger
}

// CustomTimeEncoder 自定义时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

/*
使用

	log := config.GetLogger()
	log.Info("用户登录成功")
	log.Error("数据库错误", zap.Error(err))
*/
