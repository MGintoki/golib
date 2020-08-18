package golib

import (
	"fmt"
	"time"
)

func init() {
	StartTimer(testTimer)
}

var counter = 0

func testTimer() {
	now := time.Now()
	fmt.Println(now)
	counter++
	fmt.Println("counter: %v \n", counter)
}
