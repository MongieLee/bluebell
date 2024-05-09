package logger

import (
	"bluebell/setting"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

func Init(cfg *setting.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	fmt.Println(writeSyncer)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}

func getEncoder() zapcore.Encoder {
	encoderCfg := zapcore.EncoderConfig{
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		TimeKey:        "time",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderCfg)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
