package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	//funkcija f uzima string kao parametar
	for _, v := range n {
		f(v)
	}

}

func circleArea(r float64) float64 { //obavezno navesti tip povratne promenljive float64 na kraju
	return math.Pi * r * r
}
func main() {
	// sayGreeting("andjela")
	// sayGreeting("aleksa")
	// sayBye("andjela")

	cycleNames([]string{"andjela", "ana", "marija"}, sayGreeting)
	cycleNames([]string{"andjela", "ana", "marija"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
