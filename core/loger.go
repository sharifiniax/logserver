package logger

import (
	"io"
	"log"
	"os"
	"github.com/natefinch/lumberjack"
)

var (
	logFile    *lumberjack.Logger
	logLevel   LogLevel
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

func SetLogFile(logFilePath string, maxSizeMB, maxBackups, maxAgeDays int) error {
	logFile = &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    maxSizeMB,      // حداکثر حجم فایل لاگ (MB)
		MaxBackups: maxBackups,     // حداکثر تعداد فایل‌های لاگ ذخیره شده
		MaxAge:     maxAgeDays,     // حداکثر مدت زمان نگهداری فایل‌های لاگ (روز)
			}
	
	// تنظیم خروجی لاگ بر روی فایل
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return nil
}

func SetLogLevel(level LogLevel) {
	logLevel = level
}

func Log(tag string, level LogLevel, message string) {
	if level >= logLevel {
		log.Printf("[%s] %s: %s", level.String(), tag, message)
	}
}

func CloseLogFile() {
	logFile.Close()
}

func (l LogLevel) String() string  {
	switch l{ 

	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

