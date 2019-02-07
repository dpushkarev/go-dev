package main

import (
	"sync"
	"fmt"
	"runtime"
	"pack"
)

func main(){
	pack.Echo()
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++  {
		wg.Add(1)
		go Worker(i, wg)
	}
	wg.Wait()
	//fmt.Scanln()
}

func Worker(i int, wg *sync.WaitGroup){
	defer wg.Done()
	for j := 0; j < 2 ; j++  {
		fmt.Println("num go: ", i,", step: ", j)
		runtime.Gosched()
	}
}