package utils_test

import (
	"fmt"
	"testing"

	"github.com/pro200/go-utils"
)

func TestOnlyOne(t *testing.T) {
	err := utils.OnlyOne("hello")
	fmt.Println(err)
}
