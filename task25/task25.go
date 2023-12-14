package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	fmt.Println("Before sleep")
	Sleep(5 * time.Second)
}
