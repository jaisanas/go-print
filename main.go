package mylogger

import (
	"fmt"
	"log"
)

func LogInfo(message string) {
    colorGreen := "\033[32m"
	fmt.Print(string(colorGreen))
	log.Printf("INFO - %v", message)
}
