package utils

import (
	"fmt"
	"os"
	"slices"
)

func PrintVersion(version string) {
	if args := os.Args[1:]; len(args) >= 1 {
		if slices.Contains([]string{"-v", "version"}, args[0]) {
			fmt.Println(version)
			os.Exit(0)
		}
	}
}
