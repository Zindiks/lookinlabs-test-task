package utils

import "github.com/fatih/color"

func Success(msg string) string {
	return color.New(color.FgGreen).Sprint(msg)
}

func Error(msg string) string {
	return color.New(color.FgRed).Sprint(msg)
}

func Info(msg string) string {
	return color.New(color.FgCyan).Sprint(msg)
}

func Warning(msg string) string {
	return color.New(color.FgYellow).Sprint(msg)
}
