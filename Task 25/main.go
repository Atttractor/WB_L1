package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sleep(3 * time.Second)
	fmt.Println(time.Since(start).Seconds())
}

func sleep(t time.Duration) {
	ticker := time.NewTicker(t)
	<-ticker.C
}

func sleep1(t time.Duration) {
	<-time.After(t)
}