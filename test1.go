package main

import "fmt"

func main() {
	a := []int{20, 5 , 10, 3}
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
		} else if (value < max1 ) {
			max2 = max1
			max1 = value
			value = max2 
		}

	}
	fmt.Print(max2)
	return
}
