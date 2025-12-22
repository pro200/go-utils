package utils

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/soellman/pidfile"
)

// OnlyOne은 프로세스 인스턴스가 하나만 실행되도록
func OnlyOne(name string) error {
	pidFilePath := filepath.Join("/tmp", name+".pid")
	err := pidfile.Write(pidFilePath)
	if err == nil {
		return nil
	}

	// 프로세스 이미 실행 중
	if errors.Is(err, pidfile.ErrProcessRunning) {
		return fmt.Errorf("process %q is already running", name)
	}

	// 오래되었거나 잘못된 pidfile → 정리하고 다시 시도
	// ErrFileStale:   "pidfile은 존재하지만 프로세스가 실행되고 있지 않습니다
	// ErrFileInvalid: "pidfile의 내용이 잘못되었습니다"
	if errors.Is(err, pidfile.ErrFileStale) || errors.Is(err, pidfile.ErrFileInvalid) {
		_ = pidfile.Remove(pidFilePath)
		if wErr := pidfile.Write(pidFilePath); wErr != nil {
			return fmt.Errorf("failed to rewrite pidfile: %w", wErr)
		}
		return nil
	}

	return fmt.Errorf("failed to create pidfile: %w", err)
}
