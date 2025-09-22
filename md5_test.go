package utils_test

import (
	"testing"

	"github.com/pro200/go-utils"
)

func TestMd5(t *testing.T) {
	str := utils.Md5("123456789")
	if str != "25f9e794323b453885f5181f1b624d0b" {
		t.Error("Wrong result")
	}
}
