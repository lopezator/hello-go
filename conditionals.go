package main

import "fmt"

func main() {
	//Else if
	yourAge := 18

	if yourAge >= 16  {
		fmt.Println("You can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You can have fun")
	}

	//Switch case
	switch yourAge {
		case 16 : fmt.Println("Go Drive")
		case 18 : fmt.Println("Go Vote")
		default: fmt.Println("Go Have Fun")
	}
}
