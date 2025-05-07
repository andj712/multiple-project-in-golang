package main

import "fmt"

var score = 99.5

func main() {
	//uokviru su istog paketa main, pa mozemo da pristupimo sayhello funkciiji iz greetings.go
	//oba fajla moraju biti pokrenuta, ne moze samo main.go jer onda nece moci da pozove sayhello iz greetings.go
	sayHello("andjela")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
