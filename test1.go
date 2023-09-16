package main

import "fmt"

func main() {
	a := []int{20, 20, 10, 4, 15}	
	max1 := a[0]
	max2 := 0
	if len(a) < 2 {
		return
	}
	for _, value := range a {
		if  value > max1 {
			max2 = max1
			max1 = value
			value = max2
			fmt.Println("check",max2, value, max1)
			
		} else if (value < max1 && value > max2 ) {
			fmt.Println("check",max2, value, max1)
			max2 = value
		}

	}
	fmt.Print(max2)
	return
}
