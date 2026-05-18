package utils_test

import (
	"testing"

	"github.com/pro200/go-utils"
)

func TestOnlyOne(t *testing.T) {
	cleanup, err := utils.OnlyOne("myapp")
	if err != nil {
		t.Error(err)
	}
	defer cleanup()
}
