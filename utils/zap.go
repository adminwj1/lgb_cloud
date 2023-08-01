package utils

import (
	"clouds.lgb24kcs.cn/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	var logger *zap.Logger
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core)
	zap.ReplaceGlobals(logger)
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	if _, err := os.Stat(global.APP.Configuration.Log.Root_dir); err != nil {
		_ = os.MkdirAll(global.APP.Configuration.Log.Root_dir, 0777)

	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.APP.Configuration.Log.Root_dir + "/" + global.APP.Configuration.Log.FileName,
		MaxSize:    global.APP.Configuration.Log.MaxSize,
		MaxBackups: global.APP.Configuration.Log.MaxBackup,
		MaxAge:     global.APP.Configuration.Log.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
