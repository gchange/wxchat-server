package logging

import (
	"io"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapConfig
// Level:
// EncodeLevel: LowercaseLevelEncoder(defalt), LowercaseColorLevelEncoder, CapitalLevelEncoder, CapitalColorLevelEncoder
type ZapConfig struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`                   // 级别
	Format      string `mapstructure:"format" json:"format" yaml:"format"`                // 输出
	Prefix      string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                // 日志前缀
	Dir         string `mapstructure:"dir" json:"dir"  yaml:"dir"`                        // 日志文件夹
	EncodeLevel string `mapstructure:"encodeLevel" json:"encodeLevel" yaml:"encodeLevel"` // 编码级
	LogInFile   bool   `mapstructure:"logInFile" json:"logInFile" yaml:"logInFile"`       // 输出控制台
}

func Init(zapConfig *ZapConfig) *zap.Logger {
	// log level
	logLevel := zap.NewAtomicLevel()
	switch strings.ToUpper(zapConfig.Level) {
	case "DEBUG":
		logLevel.SetLevel(zap.DebugLevel)
	case "INFO":
		logLevel.SetLevel(zap.InfoLevel)
	case "WARN", "WARNING":
		logLevel.SetLevel(zap.WarnLevel)
	case "ERROR":
		logLevel.SetLevel(zap.ErrorLevel)
	case "PANIC":
		logLevel.SetLevel(zap.PanicLevel)
	case "FATAL":
		logLevel.SetLevel(zap.FatalLevel)
	}

	var cores []zapcore.Core
	var writers []io.Writer
	// 默认输出到 console
	consoleWriter := zapcore.Lock(os.Stdout)
	cores = append(cores, getEncoderCore(zapConfig, consoleWriter, logLevel))
	writers = append(writers, consoleWriter)

	// 输出到文件
	// if zapConfig.LogInFile {
	// set default log dir
	// 	if zapConfig.Dir == "" {
	// 		zapConfig.Dir = "./log"
	// 	}
	// 	create dir if not exists
	// 	log.Printf("create %v directory\n", zapConfig.Dir)
	// 	if err := file.CreateDir(zapConfig.Dir); err != nil {
	// 		panic(fmt.Errorf("create log dir error: %s", err.Error()))
	// 	}
	// 	fileWriter := GetWriteSyncer(fmt.Sprintf("%s/server.log", zapConfig.Dir))
	// 	cores = append(cores, getEncoderCore(zapConfig, fileWriter, logLevel))
	// 	writers = append(writers, fileWriter)
	// }

	logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	// setup global default logger
	// zap.AddCallerSkip(2) 可正确打印 caller:
	//   1. ZapLogger 的 func
	//   2. logger 的 global func
	DefaultLogger.logger = logger.WithOptions(zap.AddCallerSkip(2)).Sugar()
	DefaultLogger.level = logLevel.Level()
	DefaultLogger.writer = io.MultiWriter(writers...)

	return logger
}

// getEncoderCore 获取 encoder 的 zapcore.Core
func getEncoderCore(zapConfig *ZapConfig, writer zapcore.WriteSyncer, level zapcore.LevelEnabler) (core zapcore.Core) {
	return zapcore.NewCore(getEncoder(zapConfig), writer, level)
}

// getEncoder 获取zapcore.Encoder
func getEncoder(zapConfig *ZapConfig) zapcore.Encoder {
	if zapConfig.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig(zapConfig))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(zapConfig))
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(zapConfig *ZapConfig) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	switch {
	case zapConfig.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case zapConfig.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case zapConfig.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case zapConfig.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
