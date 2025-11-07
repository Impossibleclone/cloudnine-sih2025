package log

import (
    "log"
    "os"
)

var logger *log.Logger

func Init() {
    logger = log.New(os.Stdout, "[WIPE-APP] ", log.LstdFlags|log.Lshortfile)
}

func Info(format string, v ...interface{}) {
    logger.Printf("[INFO] "+format, v...)
}

func Warn(format string, v ...interface{}) {
    logger.Printf("[WARN] "+format, v...)
}

func Error(format string, v ...interface{}) {
    logger.Printf("[ERROR] "+format, v...)
}

func Fatal(format string, v ...interface{}) {
    logger.Fatalf(format, v...)
}
