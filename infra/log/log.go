package log

import (
	"gowebstarter/infra/config"
	"log/slog"
	"os"
	"strings"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logOnce sync.Once
)

func Init() {
	logOnce.Do(initLog)
}

func debugMode() bool {
	return strings.ToLower(config.GetConf().Server.Mode) == "debug"
}

func initLog() {
	conf := config.GetConf()

	if debugMode() {
		tmpSlog := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}))
		slog.SetDefault(tmpSlog)
	} else {
		lj := &lumberjack.Logger{
			Filename:   conf.Log.Filename,
			MaxSize:    conf.Log.MaxSize,
			MaxBackups: conf.Log.MaxBackups,
			MaxAge:     conf.Log.MaxAge,
			Compress:   conf.Log.Compress,
		}
		tmpSlog := slog.New(slog.NewJSONHandler(lj, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}))
		slog.SetDefault(tmpSlog)

		slog.Info("using production mode",
			slog.Any("filename", conf.Log.Filename),
			slog.Any("maxSize", conf.Log.MaxSize),
			slog.Any("maxBackups", conf.Log.MaxBackups),
			slog.Any("maxAge", conf.Log.MaxAge),
			slog.Any("compress", conf.Log.Compress),
		)
	}
}
