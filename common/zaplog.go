package common

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	zapLogger   *zap.Logger
	AtomicLevel zap.AtomicLevel
)

func switchLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func switchLogOutput(LogOutput, LogOutputFile string) zapcore.WriteSyncer {
	switch LogOutput {
	case "file":
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   LogOutputFile, // 日志文件路径
			MaxSize:    128,           // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 30,            // 日志文件最多保存多少个备份
			MaxAge:     7,             // 文件最多保存多少天
			Compress:   true,          // 是否压缩
		})
	default:
		return zapcore.AddSync(os.Stdout)
	}
}

func Initialization(LogLevel, LogOutput, LogOutputFile string) *zap.Logger {
	if zapLogger != nil {
		zapLogger.Warn(" zapLogger already init ...")
		return zapLogger
	}

	AtomicLevel = zap.NewAtomicLevelAt(switchLogLevel(LogLevel))

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Logger",
		CallerKey:      "Caller",
		MessageKey:     "Message",
		StacktraceKey:  "Stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                  // 编码器配置
		zapcore.NewMultiWriteSyncer(switchLogOutput(LogOutput, LogOutputFile)), // 打印到控制台和文件
		AtomicLevel, // 日志级别
	)

	zapLogger = zap.New(core, zap.AddCaller(), zap.Development())

	zapLogger.Info("初始化日志", zap.String("Level", AtomicLevel.String()))

	return zapLogger
}
