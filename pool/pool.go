// Based on example found at:
// https://gist.github.com/glesica/a2cd983fa1b4400c158e

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make(chan int, 10)
	squares := make(chan int, 10)

	squareGroup := new(sync.WaitGroup)
	printGroup := new(sync.WaitGroup)

	// Create 1 worker to print the squares
	printGroup.Add(1)
	go func() {
		for square := range squares {
			fmt.Println(square)
		}
		printGroup.Done()
	}()

	// Create 4 workers to square numbers
	for i := 0; i < 4; i++ {
		squareGroup.Add(1)
		go func() {
			for number := range numbers {
				squares <- (number * number)
			}
			squareGroup.Done()
		}()
	}

	// Square some numbers!
	for i := 0; i < 10; i++ {
		numbers <- i
	}

	// Wait for all the numbers to be squared
	close(numbers)
	squareGroup.Wait()

	// Wait for all the squares to be printed
	close(squares)
	printGroup.Wait()
}
