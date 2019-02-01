package main

import (
	"fmt"
	"context"
	"time"
	"math/rand"
)

func main()  {
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int, 1)
	for i := 0; i <= 5; i++ {
		go worker(ctx, i, result)
	}
	foundBy := <- result
	fmt.Println("result", foundBy)
	finish()

	time.Sleep(time.Second)

}

func worker(context context.Context, i int, result chan int){
	waittime := time.Duration(rand.Intn(100) * 10) * time.Millisecond
	fmt.Println(i, " sleep", waittime)
	select {
	case <-context.Done():
		return
	case <-time.After(waittime):
		fmt.Println("worker ", i, "done")
		result<-i
	}
}