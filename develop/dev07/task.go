package main

import (
	"fmt"
	"time"
)

func or(inChannels ...<-chan int) chan int {
	ch := make(chan int)
	closeChan := make(chan struct{})
	for _, c := range inChannels {
		go func(c <-chan int) {
			for v := range c {
				ch <- v
			}
			closeChan <- struct{}{}
		}(c)
	}
	go func() {
		defer close(ch)
		count := 0
		for range closeChan {
			count++
			if count == len(inChannels) {
				return
			}
		}
	}()
	return ch
}

func main() {
	channel1 := make(chan int, 1024)
	go func() {
		defer close(channel1)
		for i := 0; i < 5; i++ {
			channel1 <- i
			time.Sleep(time.Millisecond * 500)
		}
	}()
	channel2 := make(chan int, 1024)
	go func() {
		defer close(channel2)
		for i := 0; i < 10; i++ {
			channel2 <- i * 5
			time.Sleep(time.Millisecond * 500)
		}
	}()
	result := or(channel1, channel2)
	for value := range result {
		fmt.Println(value)
	}
}
