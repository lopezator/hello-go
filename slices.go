package main

import "fmt"

func main() {
	numSlice := []int {5, 4, 3, 2, 1, 0}
	numSlice2 := numSlice[3:5]

	fmt.Println("numSlice2[0] = ", numSlice2[0])
	fmt.Println("numSlice2[1] = ", numSlice2[1])
	fmt.Println("numSlice[:2] = ", numSlice[:2])
	fmt.Println("numSlice[2:] = ", numSlice[2:])

	//Los 5 primeros, con valor 0, tamaño maximo de 10
	numSlice3 := make([]int, 6, 10)

	//Copia los valores de numSlice a numSlice3, pero solo hasta el index 5
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3[0])

	//Añade valores al numSlice
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3[6])
}