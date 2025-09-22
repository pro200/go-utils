package utils_test

import (
	"testing"

	"github.com/pro200/go-utils"
)

func TestUid(t *testing.T) {
	str := utils.Base62(123456789)
	if str != "8m0Kx" {
		t.Error("Wrong result")
	}
}
