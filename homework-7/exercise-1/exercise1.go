package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start at " + time.Now().Format("15:04:05\n\r"))
	go spinner(50 * time.Millisecond)
	duration := time.Second * 10
	time.Sleep(duration)
	fmt.Printf("End at " + time.Now().Format("15:04:05\n\r"))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
