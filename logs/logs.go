package logs

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/byteweap/pkg/filex"
	"github.com/rs/zerolog"
)

var logger *Loggerx

type Loggerx struct {
	zerolog.Logger
	mux          sync.Mutex
	interval     int       // 日志切割时间间隔, 单位:h
	lastFileTime time.Time // 上次log文件创建时间
	path         string    // 日志文件存放路径
}

func Init(level, pathname string, interval int) {

	switch strings.ToLower(level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel) // 默认debug级别
	}

	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix // 更快更小
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.TimestampFieldName = "Time"
	zerolog.LevelFieldName = "Level"
	zerolog.MessageFieldName = "Msg"

	logger = &Loggerx{
		Logger:       zerolog.New(newOutput(pathname)).With().Logger(),
		interval:     interval,
		mux:          sync.Mutex{},
		lastFileTime: time.Now(),
		path:         pathname,
	}
}

// 获取输出 控制台/文件
// 当pathname为空时,输出到控制台
func newOutput(pathname string) io.Writer {

	// 1. 默认标准输出
	// 2. 文件夹设置不为空时,写入文件
	if pathname != "" {
		now := time.Now().Format("2006-01-02_15:04:05")
		filename := fmt.Sprintf("%s.log", now)

		// 文件夹不存在,则创建
		if !filex.IsExist(pathname) {
			err := os.MkdirAll(pathname, os.ModePerm)
			if err != nil {
				fmt.Println("MkdirAll path[", pathname, "] error:", err.Error())
			}
		}
		file, err := os.Create(path.Join(pathname, filename))
		if err == nil {
			return file
		} else {
			fmt.Println("create file[", filename, "] error:", err.Error())
		}
	}
	return os.Stdout
}

// Debug 适配旧的logs打印
func Debug(format string, v ...interface{}) {
	Debugx().Msgf(format, v...)
}

func Info(format string, v ...interface{}) {
	Infox().Msgf(format, v...)
}

func Warn(format string, v ...interface{}) {
	Warnx().Msgf(format, v...)
}

func Error(format string, v ...interface{}) {
	Errorx().Msgf(format, v...)
}

// Fatal 打印Fatal信息 (程序终止)
func Fatal(format string, v ...interface{}) {
	Fatalx().Msgf(format, v...)
}

// Panic 打印Panic信息 (程序不终止)
func Panic(format string, v ...interface{}) {
	Panicx().Msgf(format, v...)
}

func Custom(tag, format string, v ...interface{}) {
	log().Any("Tag", tag).Msgf(format, v...)
}

// ----------------- 链式操作 log ------------

func Debugx() *zerolog.Event {
	return newEvent(zerolog.DebugLevel)
}

func Infox() *zerolog.Event {
	return newEvent(zerolog.InfoLevel)
}

func Errorx() *zerolog.Event {
	return newEvent(zerolog.ErrorLevel)
}

func Warnx() *zerolog.Event {
	return newEvent(zerolog.WarnLevel)
}

// Fatalx Fatal消息打印 (程序终止)
func Fatalx() *zerolog.Event {
	return newEvent(zerolog.FatalLevel)
}

// Panicx Panic消息打印 (程序不会终止)
func Panicx() *zerolog.Event {
	return newEvent(zerolog.PanicLevel)
}

func log() *zerolog.Event {
	return newEvent(zerolog.NoLevel)
}

func newEvent(level zerolog.Level) *zerolog.Event {
	// 1.检测logger引擎是否初始化、是否要切割
	logger.check()
	// 2.根据level返回Event
	switch level {
	case zerolog.DebugLevel:
		return logger.Logger.Debug().Timestamp()
	case zerolog.InfoLevel:
		return logger.Logger.Info().Timestamp()
	case zerolog.WarnLevel:
		return logger.Logger.Warn().Timestamp()
	case zerolog.ErrorLevel:
		return logger.Logger.Error().Timestamp()
	case zerolog.FatalLevel:
		return logger.Logger.Fatal().Timestamp()
	case zerolog.PanicLevel:
		return logger.Logger.Panic().Timestamp()
	case zerolog.NoLevel:
		return logger.Logger.Log().Timestamp()
	default:
		return logger.Logger.Debug().Timestamp()
	}
}

func (log *Loggerx) check() {
	if log == nil {
		Init("debug", "", 0)
	} else {
		// 日志文件切割
		if log.interval > 0 && time.Now().Add(-time.Hour*time.Duration(log.interval)).After(log.lastFileTime) {
			log.mux.Lock()
			Init("debug", log.path, log.interval)
			log.mux.Unlock()
		}
	}
}
