package main

import (
	"fmt"
	"math/rand"
)

func Generate(chNumber chan <- int, times int) {
	for i := 0; i < times; i++ {
		number := rand.Intn(100)
		chNumber <- number
		fmt.Printf("Generated number %d\n", number)
	}
	close(chNumber)
}

func Increment(chNumber <- chan int, chIncremented chan <- int) {
	for number := range chNumber {
		incrementedNumber := number + 1
		chIncremented <- incrementedNumber
		fmt.Printf("Incremented number %d\n", incrementedNumber)
	}
	close(chIncremented)
}

func main() {
	chNumber := make(chan int, 5)
	chIncremented := make(chan int, 5)
	go Generate(chNumber, 10)
	go Increment(chNumber, chIncremented)
	
	for number := range chIncremented {
		fmt.Printf("Number incremented is %d\n", number)
	}
}
