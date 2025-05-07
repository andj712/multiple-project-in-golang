package main

import (
	"fmt"
	"strings"
)

func main() {

	greeting := "hello world"
	//vraca se true ili false ako sadrzi taj string
	//fmt.Println(strings.Contains(greeting, "hello"))
	//fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	//fmt.Println(strings.ToUpper(greeting))//velika slova
	//fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " ")) //razdvaja po razmaki

	// fmt.Println("original string value= ", greeting)

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort.Ints(ages) //sortira niz
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30) //vraca poziciju gde se 30 nalazi, ali ovde u sortiranom nizu jer smo ga sortirali
	// fmt.Println(index)

	// names := []string{"andjela", "djordje", "elena"}
	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "djordje"))
	//LOOPS
	// x := 0 //koristi se for uvek, ne while
	// for x < 5 {
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is: ", i)

	// }

	names := []string{"marko", "petar", "janko"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
		value = "new string" //ovo ne menja value unutar name, jer je kopija samo
	}

	// for _, value := range names { //ovako _ umesto indexa kada ne zelimo da koristimo/ispisemo index da nam ne bi bacalo gresku
	// 	fmt.Printf("the value is %v \n", value)
	// }

	fmt.Println(names)

}
