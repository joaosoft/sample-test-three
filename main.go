package main

import (
	"fmt"

	logger "github.com/joaosoft/logger"
)

func main() {
	fmt.Println(NewSampleTestThree())
}

func NewSampleTestThree() string {
	logger.Info("executing info logger for sample-test-three TAGGED")
	return "hello, i'm the sample-test-three TAGGED"
}
