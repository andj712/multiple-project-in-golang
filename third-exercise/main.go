package main

import (
	"fmt"
)

func main() {
	age := 45

	fmt.Println(age <= 50) //vraca true ili false
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	names := []string{"marko", "luka", "ana", "petar", "sara"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue //izlazi iz trenutne iteracije for loop i prelazi na sledecu
		}
		if index > 2 {
			fmt.Println("breaking at position", index)
			break //izlazi iz cele for petlje
		}
		fmt.Printf("value at pos %v is %v", index, value)
	}
}
