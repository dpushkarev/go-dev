package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(time.Minute)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step: ", i, "time: ", tickTime)
		if i == 5 {
			fmt.Println("stop ticker")
			ticker.Stop()
			break
		}
	}
}