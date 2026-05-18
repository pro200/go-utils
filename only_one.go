package utils

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/soellman/pidfile"
)

// OnlyOne은 프로세스 인스턴스가 하나만 실행되도록
func OnlyOne(name string) (cleanup func(), err error) {
	pidFilePath := filepath.Join("/tmp", name+".pid")

	write := func() error { return pidfile.Write(pidFilePath) }

	err = write()
	if err != nil {
		switch {
		case errors.Is(err, pidfile.ErrProcessRunning):
			return nil, fmt.Errorf("process %q is already running", name)
		case errors.Is(err, pidfile.ErrFileStale), errors.Is(err, pidfile.ErrFileInvalid):
			_ = pidfile.Remove(pidFilePath)
			if err = write(); err != nil {
				return nil, fmt.Errorf("failed to rewrite pidfile: %w", err)
			}
		default:
			return nil, fmt.Errorf("failed to create pidfile: %w", err)
		}
	}

	cleanup = func() { _ = pidfile.Remove(pidFilePath) }
	return cleanup, nil
}
