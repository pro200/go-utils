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

func Exit(v ...any) {
	fmt.Print(v...)
	os.Exit(0)
}

func Exitf(format string, v ...any) {
	fmt.Printf(format, v...)
	os.Exit(0)
}

func Input(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
