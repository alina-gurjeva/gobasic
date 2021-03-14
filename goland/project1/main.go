package main

import (
	"fmt"

	"./tosay"

	"./viar"
)

func main() {
	fmt.Println("Let's GO!")
	fmt.Println(tosay.Isay())
	a := 5
	b := "stringsome"
	fmt.Printf("%T - %v, %T, %v\n", a, a, b, b)
	cat := tosay.Catsay()
	dog := tosay.Dogsay()
	fmt.Printf("%v - it's cat, %v - it's dog", cat, dog)
	fmt.Print(viar.Visible)

}
