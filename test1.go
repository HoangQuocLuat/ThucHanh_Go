package main

import "fmt"

func main() {
	a := []int{5, 5, 2, 6}
	if len(a) < 2 {
		return
	}
	max1 := a[0]
	max2 := 0
	for _, value := range a {
		if  value > max1 {
			max2 = max1
			max1 = value
			value = max2
		} 

	}
	fmt.Print(max2)
	return
}
