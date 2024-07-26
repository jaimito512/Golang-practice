package main

import "fmt"

type digit []int

func main() {

	b := []int{}
	for i := 0; i < 11; i += 1 {
		b = append(b, i)
	}
	fmt.Println(b)

	for j := 1; j < 11; j += 1 {
		if b[j]%2 == 0 {
			fmt.Println(j, "is even")
		} else {
			fmt.Println(j, "is odd")
		}

	}
}
