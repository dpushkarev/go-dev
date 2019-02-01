package main

import (
	"fmt"
)

func main(){
	stopCh := make(chan struct{})
	ch := make(chan int)

	go func(stopCh chan struct{}, ch chan int) {
		val := 0
		for {
			select{
			case <- stopCh:
				return
			case ch <- val:
				val++
			}
		}

	}(stopCh, ch)

	for item := range ch {
		fmt.Println("value from ch: ", item)
		if item >= 3 {
			fmt.Println("put in stopCh")
			stopCh <- struct{}{}
			break
		}
	}
}