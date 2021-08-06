/**
 * @Author: sickle
 * @Date: 2021/8/6
 */
package utils

import (
	conf "fengcheqiu/configs"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger

func init() {

	writer := getLogWriter(
		conf.Conf.Zap.File,
		conf.Conf.Zap.MaxSize,
		conf.Conf.Zap.MaxBackups,
		conf.Conf.Zap.MaxAge,
		conf.Conf.Zap.Compress,
	)
	encoder := getEncoder()

	var allCore []zapcore.Core

	level := zap.DebugLevel

	switch conf.Conf.App.Level {
	case "debug":
		level = zap.DebugLevel
		consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
		allCore = append(allCore, consoleCore)
	case "info":
		level = zap.InfoLevel
	default:
		level = zap.InfoLevel
	}

	fileCore := zapcore.NewCore(encoder, writer, level)

	allCore = append(allCore, fileCore)
	core := zapcore.NewTee(allCore...)

	Logger = zap.New(core, zap.AddCaller())

	return
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	defer logger.Close()
	return zapcore.AddSync(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
