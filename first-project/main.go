package main

import "fmt"

func main() {

	//STRINGs
	// var nameOne string = "mario"
	// var nameTwo = "luigi"
	// var nameThree string //definisali smo string i mora biti string ne sme se promeniti

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "yoshi" //ovo samo unutar funkcije moze ovako da se deklarise
	// fmt.Println(nameFour)

	//ints

	// var ageOne int = 23
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	//int8- od -128 do 127
	//int 16 od -32.768 do 32767
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// //unsigned int znaci da nema negativnog broja, ne bi mogli da stavimo da je -25
	// //uint8 od 0 do 256, povecano je u odnosu na int8 jer nemamo negativne brojeve
	// var numThree uint = 25

	// var scoreOne float32=25.98
	// var scoreTwo float64= 88885858585858.7
	// // scoreThree:=1.5 //postaje float64 i ne moze da se menja kasnije
	// name := "Andjela"
	// fmt.Println("hello, my age is ", ageOne)

	//formatted string, %_ = format specifier
	// fmt.Printf("my age is %v and my name is %v", ageOne, name) //%v kao varijabla
	// fmt.Printf("my age is %q and my name is %q", ageOne, name) //%q koristi se nad stringovima, stavi "" navodnike, ne moze nad int
	// fmt.Printf("age is of type %T", ageOne)                    //%T vraca tip varijable
	// fmt.Printf("you scored %f points", 225.55)
	// fmt.Printf("you scored %0.1f points", 225.55) //limitira ima zareza koliko je

	//sprint f save formatted string
	//sacuva nam u varijablu
	//var str = fmt.Sprintf("my age is %v and my name is %v", ageOne, name)
	//fmt.Println("the saved string is: ", str)

	//var ages [3]int=[3]int {20,25,30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"andjela", "aleksandra", "mima", "aleksa"}
	names[2] = "ana"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//SLICES(koriste slicno kao nizovi jer nema definisan broj elemenata)
	//kod njih mozemo da ubacimo novi element

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slice ranges
	//krece se od nulte pozicije ne prve
	rangeOne := names[1:3] //vraca prvu i drugu poziciju ali ne i trecu, zapravo cetvrtu kad racunamo da je nula prvi el
	rangeTwo := names[2:3]
	rangeThree := names[:3] //od pocetka do trece
	rangeFour := names[2:]  //od druge pozicije do kraja

	fmt.Println(rangeOne, rangeTwo, rangeThree, rangeFour)

}
