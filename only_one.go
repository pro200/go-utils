package utils

import (
	"errors"
	"os"
	"strings"

	"github.com/soellman/pidfile"
)

func OnlyOne() error {
	paths := strings.Split(os.Args[0], "/")
	fileName := paths[len(paths)-1]

	if fileName == "main" {
		return errors.New("main file is not allowed.")
	}

	pidFilePath := "/tmp/" + fileName + ".pid"
	err := pidfile.Write(pidFilePath)

	if errors.Is(err, pidfile.ErrProcessRunning) {
		return errors.New("process is already running.")
	}

	// rewrite pidfile
	// ErrFileStale:   "pidfile은 존재하지만 프로세스가 실행되고 있지 않습니다
	// ErrFileInvalid: "pidfile의 내용이 잘못되었습니다"
	if errors.Is(err, pidfile.ErrFileStale) || errors.Is(err, pidfile.ErrFileInvalid) {
		pidfile.Remove(pidFilePath)
		pidfile.Write(pidFilePath)
		return nil
	}

	return err
}
