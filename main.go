package logger

import (
	"fmt"
)

func LogInfo(message string) {
    colorGreen := "\033[32m"
	colorReset := "\033[0m"
	info := fmt.Sprintf("INFO -  %v", message)
	fmt.Println(string(colorGreen), info, string(colorReset))
}

func LogWarn(message string) {
    colorGreen := "\033[32m"
	colorReset := "\033[0m"
	info := fmt.Sprintf("WARN -  %v", message)
	fmt.Println(string(colorGreen), info, string(colorReset))
}

func LogError(message string) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	info := fmt.Sprintf("ERROR -  %v", message)
	fmt.Println(string(colorRed), info, string(colorReset))
}