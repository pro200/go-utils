package utils_test

import (
	"testing"

	"github.com/pro200/go-utils"
)

func TestStrPad(t *testing.T) {
	result := utils.StrPad("1345", 6, "0")
	if result != "001345" {
		t.Error("Wrong result")
	}
}
