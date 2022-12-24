package main

import "fmt"

func main() {
	a := [...]int{10,20,30,40,50,60}
	// b := a[1:4]
	// fmt.Println(b) // [20 30 40]
	// a[2] = 0
	// fmt.Println(a)
	// fmt.Println(b)
	for key := range a{
		fmt.Println(a[key])
	}
}