package log

import (
	"fmt"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func InitLogger(appName, logPath, level string) error {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	logWriter, err := rotatelogs.New(
		logPath+string(filepath.Separator)+appName+".log.%Y_%m_%d_%H",
		rotatelogs.WithLinkName(logPath+string(filepath.Separator)+appName+".log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
	)
	if err != nil {
		return fmt.Errorf("failed to create rotatelogs: %w", err)
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(logWriter), levelMap[level])

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	return nil
}

func Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

func DPanicf(msg string, args ...interface{}) {
	logger.DPanicf(msg, args...)
}

func Panicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}
