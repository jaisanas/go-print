package logger

import (
	"fmt"
	"time"
)

func LogInfo(message string) {
    colorGreen := "\033[32m"
	colorReset := "\033[0m"
	info := fmt.Sprintf("%v INFO -  %v", time.Now(), message)
	fmt.Println(string(colorGreen), info, string(colorReset))
}

func LogWarn(message string) {
    colorYellow := "\033[33m"
	colorReset := "\033[0m"
	info := fmt.Sprintf("%v WARN -  %v", time.Now(), message)
	fmt.Println(string(colorYellow), info, string(colorReset))
}

func LogError(message string) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	info := fmt.Sprintf("%v ERROR -  %v", time.Now(), message)
	fmt.Println(string(colorRed), info, string(colorReset))
}