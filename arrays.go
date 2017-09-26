package main

import "fmt"

func main() {
	var favNums[5] float64

	favNums[0] = 2
	favNums[1] = 15
	favNums[2] = 99.45
	favNums[3] = 75
	favNums[4] = 77

	fmt.Println(favNums[2])

	favNums2 := [5]float64 {1, 2, 3, 4, 5}

	for i, value := range favNums2 {
		fmt.Println(value, i)
	}
}
