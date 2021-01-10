package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}
	for i, id := range ids {
		fmt.Printf("%d - ID:%d\n", i, id)
	}
	for _, id := range ids {
		fmt.Printf("ID:%d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)
	//range with map
	emails := map[string]string{"aman": "amanlalwani0807@gmail.com", "aman1": "lalwaniaman0807@gmail.com", "aman2": "amanlalwani0807@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Name: " + k)
	}

}
