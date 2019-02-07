package main

import(
	"fmt"
	"time"
)


func getComments() chan string{
	result := make(chan string, 1)
	go func(out chan string){
		time.Sleep(2 * time.Second)
		fmt.Println("get comments")
		out <- "32 comments"
	}(result)
	return result
}

func getPages(){
	chitems := getComments()
	time.Sleep(1 * time.Second)
	fmt.Println("get news")

	item := <-chitems
	fmt.Println("comments:", item)

}

func main()  {
	for i := 0; i < 1; i++{
		getPages()
	}
}