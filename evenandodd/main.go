package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, _s := range s {
		p := "odd"
		if _s%2 == 0 {
			p = "even"
		}

		fmt.Printf("%v is %v\n", _s, p)
	}
}
