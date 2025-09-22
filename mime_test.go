package utils_test

import (
	"github.com/pro200/go-utils"
	"testing"
)

func TestMime(t *testing.T) {
	ct := utils.ContentType("dog.jpg")
	if ct != "image/jpeg" {
		t.Error("Wrong result")
	}
}
