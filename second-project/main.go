package main

import (
	"fmt"
	"sort"
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

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) //sortira niz
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) //vraca poziciju gde se 30 nalazi, ali ovde u sortiranom nizu jer smo ga sortirali
	fmt.Println(index)

	names := []string{"andjela", "djordje", "elena"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "djordje"))
}
