package utils_test

import (
	"testing"

	"github.com/pro200/go-utils"
)

func TestMime(t *testing.T) {
	ct := utils.ContentType("dog.html")
	if ct != "text/html" {
		t.Error("Wrong result")
	}
}
