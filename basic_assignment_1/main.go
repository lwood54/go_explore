package main

import "fmt"

func main() {
	sInts := []int{}
	for i := 0; i <= 10; i++ {
		sInts = append(sInts, i)
	}

	for _, num := range sInts {
		if num%2 == 0 {
			fmt.Println("even: ", num)
		} else {
			fmt.Println("odd: ", num)
		}
	}
}
