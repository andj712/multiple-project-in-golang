package main

import (
	"fmt"
)

func main() {

	menu := map[string]float64{
		//key-string a value je float
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
		"coffe": 2.19,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"]) //mora u navodnicima vraca cenu od pie

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		06123221: "andjela",
		06231212: "marija",
		06346434: "ana",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[06123221]) //vraca value tj andjela za ovaj broj

	phonebook[06123221] = "aleksa"
	fmt.Println(phonebook)

	phonebook[06346434] = "marko"
	fmt.Println(phonebook)

}
