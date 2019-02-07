package main

//ограничение кол-во одновременно работающих горутин при помощи канала

import (
	"sync"
	"fmt"
	"runtime"
)

const(
	quoteLinmit = 2
	iterates = 5
)

func main()  {
	qouteCh := make(chan struct{},quoteLinmit)
	wg := &sync.WaitGroup{}
	for i := 0; i < iterates; i++ {
		wg.Add(1)
		go worker(qouteCh, i, wg)
	}
	wg.Wait()
}

func worker(qoutech chan struct{}, num int, wg *sync.WaitGroup){
	defer wg.Done()
	qoutech <- struct{}{}
	for i := 0; i < iterates; i++ {
		fmt.Println("num go: ", num, ", step: ", i)
		if i % 2 == 0{
			<-qoutech
			qoutech <- struct{}{}
		}
		runtime.Gosched()
	}
	<-qoutech
}
