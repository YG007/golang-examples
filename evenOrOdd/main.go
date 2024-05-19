package main

import "fmt"

func main() {
	values := []int{}

	for i := 0; i <= 10; i += 1 {
		values = append(values, i)
	}

	for _, val := range values {
		ans := "Odd"
		if val%2 == 0 {
			ans = "Even"
		}

		fmt.Println(val, ans)
	}
}
