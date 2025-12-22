package utils_test

import (
	"testing"

	"github.com/pro200/go-utils"
)

func TestLogger(t *testing.T) {
	logger, err := utils.NewLogger("hello", "/tmp/test", 10)
	if err != nil {
		t.Error(err)
	}
	defer logger.Close()

	logger.Info("hello")
}
