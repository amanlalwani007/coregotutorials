package main

import "fmt"

func main() {
	x := 5
	y := 10
	if x < y {
		//%d is used for base 10 int
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "blk"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "black" {
		fmt.Println("color is black")
	} else {
		fmt.Println("color is not red or not blue")
	}

	switch color {
	case "red":
		fmt.Println("color is red ")
	case "black":
		fmt.Println("color is black")
	default:
		fmt.Println("color is not red and not black")

	}
}
