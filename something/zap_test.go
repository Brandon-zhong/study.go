package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

//测试实例来着 https://www.liwenzhou.com/posts/Go/zap/
//基础zap使用测试
func TestZapBase(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("fail to fetch url",
		zap.String("url", "www.baidu.com"),
		zap.Int("ttl", 20),
		zap.Duration("offline", time.Second),
	)
}
//zap写文件相关测试
func TestZapWriteFile(t *testing.T) {
	filePath := "/data/log/info.log"

	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})
	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	syncer, fileClose, err := zap.Open(filePath)
	if err != nil {
		fileClose()
		return
	}
	hishCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), syncer, highPriority)
	lowCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), syncer, lowPriority)

	logger := zap.New(zapcore.NewTee(hishCore, lowCore), zap.AddCaller())
	for i := 0; i < 1000; i++ {
		logger.Info("this is info level", zap.String("hahah", "string info"))
	}
	logger.Sync()
}

//zap 集成分隔文件库-lumberjack
func TestZapSplitFile(t *testing.T) {
	//生成日志文本同步器
	lumberJackLogger := lumberjack.Logger{
		Filename:   "/data/log/info.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	syncer := zapcore.AddSync(&lumberJackLogger)

	//生成日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, syncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	for i := 0; i < 100000; i++ {
		logger.Info("this is lumberjack demo",
			zap.String("hahah", "hahah"),
			zap.Int("this is int value", 203910),
		)
	}
	logger.Sync()
}
