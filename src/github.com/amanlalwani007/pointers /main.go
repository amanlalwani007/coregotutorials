package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)
	//using to read from address
	fmt.Println(*b)
	//change value from ptr
	*b = 10
	fmt.Println(a)

}
