package utils_test

import (
	"testing"

	"github.com/pro200/go-utils"
)

func TestMime(t *testing.T) {
	ct := utils.ContentType("dog.jpg")
	if ct != "image/jpeg" {
		t.Error("Wrong result")
	}
}
