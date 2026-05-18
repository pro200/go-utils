package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/soellman/pidfile"
)

// OnlyOne은 프로세스 인스턴스가 하나만 실행되도록 보장한다.
// 반환된 cleanup 함수는 프로세스 종료 시 호출하여 pidfile을 정리해야 한다.
//
// 사용 예:
//
//	cleanup, err := utils.OnlyOne("myapp")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer cleanup()
func OnlyOne(name string) (cleanup func(), err error) {
	// path traversal 방어
	if name == "" || strings.ContainsAny(name, `/\`) {
		return nil, fmt.Errorf("invalid name: %q", name)
	}

	// /tmp 하드코딩 대신 OS 임시 디렉토리 사용 (TMPDIR 등 환경변수 존중)
	pidFilePath := filepath.Join(os.TempDir(), name+".pid")

	write := func() error { return pidfile.Write(pidFilePath) }

	err = write()
	if err != nil {
		switch {
		case errors.Is(err, pidfile.ErrProcessRunning):
			// 이미 실행 중인 프로세스가 있음
			return nil, fmt.Errorf("process %q is already running", name)

		case errors.Is(err, pidfile.ErrFileStale), errors.Is(err, pidfile.ErrFileInvalid):
			// stale: pidfile은 있지만 프로세스가 죽은 상태
			// invalid: pidfile 내용이 깨진 상태
			// → 정리하고 재시도
			_ = pidfile.Remove(pidFilePath)
			if err = write(); err != nil {
				return nil, fmt.Errorf("failed to rewrite pidfile at %s: %w", pidFilePath, err)
			}

		default:
			return nil, fmt.Errorf("failed to create pidfile at %s: %w", pidFilePath, err)
		}
	}

	cleanup = func() {
		// 정리 실패는 치명적이지 않지만, 다음 실행 시 stale로 잡히게 됨
		_ = pidfile.Remove(pidFilePath)
	}
	return cleanup, nil
}
