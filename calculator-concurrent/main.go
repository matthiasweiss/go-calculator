package main

import "fmt"

func main() {
	numbers := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
	}

	c := make(chan int)
	for _, v := range numbers {
		go sum(c, v)
	}

	sums := []int{}
	for range numbers {
		sums = append(sums, <-c)
	}

	fmt.Println(sums)
}

func sum(c chan<- int, numbers []int) {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	c <- sum
}
