package helpers

import "fmt"

func LogInfo(message string, args ...interface{}) {
	fmt.Printf(message+"\n", args...)
}
