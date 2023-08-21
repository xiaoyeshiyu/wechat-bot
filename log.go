package main

import (
	"os"

	"wechat-bot/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Level zapcore.Level

// DebugLevel logs are typically voluminous, and are usually disabled in
// production.
//DebugLevel Level = iota - 1
// InfoLevel is the default logging priority.
//InfoLevel
// WarnLevel logs are more important than Info, but don't need individual
// human review.
//WarnLevel
// ErrorLevel logs are high-priority. If an application is running smoothly,
// it shouldn't generate any error-level logs.
//ErrorLevel

func NewLog() *zap.SugaredLogger {
	switch config.GetLogLevel() {
	case "DEBUG", "debug", "5":
		Level = zapcore.DebugLevel
	case "INFO", "info", "4":
		Level = zapcore.InfoLevel
	case "WARN", "warn", "3":
		Level = zapcore.WarnLevel
	case "ERROR", "error", "2":
		Level = zapcore.ErrorLevel
	default:
		Level = zapcore.InfoLevel
	}

	logger := NewLogger(config.GetLogFile(), Level, config.GetLogMaxSize(), config.GetLogBackups(), config.GetLogMaxAge(), true, "wechat-bot").Sugar()
	zap.ReplaceGlobals(logger.Desugar())
	return logger
}

// NewLogger
//
//	获取日志
//	filePath 日志文件路径
//	level 日志级别
//	maxSize 每个日志文件保存的最大尺寸 单位：M
//	maxBackups 日志文件最多保存多少个备份
//	maxAge 文件最多保存多少天
//	compress 是否压缩
//	serviceName 服务名
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

/**
 * zapcore构造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 包路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		// 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}
