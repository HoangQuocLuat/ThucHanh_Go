package main

import "fmt"

func main() {
	arrNumbers := []int{1, 3, 4, 5, 5, 6, 6, 6, 7, 9, 9, 8, 9, 9}

	count := make(map[int]int)

	for _, num := range arrNumbers {
		count[num]++
	}	
	fmt.Print(count)

	max := 0
	n := 0

	for index, value := range count {
		if value > max {
			max = value
			n = index
		}
	}

	fmt.Print(n, " lap lai nhieu nhat voi so lap la ", max)
}
