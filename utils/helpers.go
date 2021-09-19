package utils

import (
	"fmt"
	"os"
)

func DebugLog(message ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(message...)
	}
}
