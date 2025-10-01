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

func Input(v ...any) string {
	var input string
	fmt.Print(v...)
	fmt.Scanln(&input)
	return input
}

func Inputf(format string, v ...any) string {
	var input string
	fmt.Printf(format, v...)
	fmt.Scanln(&input)
	return input
}
