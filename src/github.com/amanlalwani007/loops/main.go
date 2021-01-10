package main

import "fmt"

func main() {
	//long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1

	}
	//short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}
	//fizz method
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

}
