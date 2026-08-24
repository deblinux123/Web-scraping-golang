package main

import "fmt"

func scrape(url string, ch chan string) {
	result := "finished:" + url

	ch <- result
}

func add(a, b int, ch chan int) {
	sum := a + b

	ch <- sum
}

func main() {
	ch := make(chan string)

	go scrape("https://huggingface.co/learn/llm-course/chapter6/3b", ch)

	result := <-ch

	fmt.Println(result)

	ch2 := make(chan int)

	go add(12, 12, ch2)

	sum := <-ch2
	fmt.Println(sum)
}
