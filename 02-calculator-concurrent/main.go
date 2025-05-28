package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	numbers := [][]int{
		{0, 1, 2},
		{4, 5, 6},
		{7, 8, 9},
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

	// simulate heavy calculation
	time.Sleep(time.Duration(rand.IntN(2)) * time.Second)

	c <- sum
}
