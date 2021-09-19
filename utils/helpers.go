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

// Don't do this in real production code
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
