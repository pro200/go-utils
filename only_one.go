package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/soellman/pidfile"
)

// OnlyOne은 프로세스 인스턴스가 하나만 실행되도록 합니다.
// /tmp(또는 사용자 지정 경로)에 pidfile을 생성하고 다른 인스턴스가 있으면 오류를 반환합니다.
func OnlyOne(pidDir ...string) error {
	fileName := filepath.Base(os.Args[0])

	// `go run`으로 직접 실행하는 것을 방지합니다.
	if fileName == "main" {
		return errors.New("running with `go run` is not supported, please build the binary first")
	}

	// 기본 pid 디렉토리: /tmp
	dir := "/tmp"
	if len(pidDir) > 0 && pidDir[0] != "" {
		dir = pidDir[0]
	}

	pidFilePath := filepath.Join(dir, fileName+".pid")
	err := pidfile.Write(pidFilePath)
	if err == nil {
		return nil
	}

	// 프로세스 이미 실행 중
	if errors.Is(err, pidfile.ErrProcessRunning) {
		return fmt.Errorf("process %q is already running", fileName)
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
