package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Logger struct {
	warn  *log.Logger
	info  *log.Logger
	error *log.Logger
	fpLog *os.File
}

func NewLogger(name, dir string, maxLogSize int) (*Logger, error) {
	var logger = new(Logger)
	var err error

	// 디렉토리 생성 (없으면)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return logger, fmt.Errorf("log dir create failed: %w", err)
	}

	// 오늘 날짜 로그 파일 생성
	today := time.Now().Format("2006-01-02")
	logFileName := name + "-" + today + ".log"
	logFilePath := filepath.Join(dir, logFileName)

	// 로그파일 오픈
	logger.fpLog, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return logger, err
	}

	logger.info = log.New(logger.fpLog, "INFO: ", log.Ldate|log.Ltime)
	logger.warn = log.New(logger.fpLog, "WARN: ", log.Ldate|log.Ltime)
	logger.error = log.New(logger.fpLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	cleanupOldLogs(dir, name, maxLogSize)
	return logger, nil
}

func (l *Logger) Info(message string) error {
	if l.info == nil {
		return fmt.Errorf("declare logger before using Info")
	}
	l.info.Println(message)
	return nil
}

func (l *Logger) Warn(message string) error {
	if l.info == nil {
		return fmt.Errorf("declare logger before using Warn")
	}
	l.warn.Println(message)
	return nil
}

func (l *Logger) Error(message string) error {
	if l.info == nil {
		return fmt.Errorf("declare logger before using Error")
	}
	l.error.Println(message)
	return nil
}

func (l *Logger) Close() {
	if l.fpLog != nil {
		l.fpLog.Close()
	}
}

func cleanupOldLogs(dir, prefix string, maxLog int) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var logFiles []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		// name-YYYY-MM-DD.log 형식만 대상
		if strings.HasPrefix(entry.Name(), prefix+"-") && filepath.Ext(entry.Name()) == ".log" {
			logFiles = append(logFiles, entry.Name())
		}
	}

	// 파일명이 날짜이므로 문자열 정렬 = 시간순 정렬
	sort.Strings(logFiles)

	// 초과분 삭제
	if len(logFiles) > maxLog {
		removeCount := len(logFiles) - maxLog
		for i := 0; i < removeCount; i++ {
			path := filepath.Join(dir, logFiles[i])
			if err := os.Remove(path); err != nil {
				return err
			}
		}
	}

	return nil
}
