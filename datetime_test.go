package utils_test

import (
	"fmt"
	"testing"

	"github.com/pro200/go-utils"
)

func TestDatetime(t *testing.T) {
	_, err := utils.ParseTime("20251029122304")
	if err != nil {
		t.Error(err)
	}

	str, err := utils.ParseTimeFormat("20251029122304", "%Y-%m-%d %H:%M:%S")
	fmt.Println("~~>", str)
	if err != nil || str != "2025-10-29 12:23:04" {
		t.Error("Wrong result")
	}
}
