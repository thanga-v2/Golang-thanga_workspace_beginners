package main

import "fmt"

type FSS struct {
	employee string
	M1 string
	M2 string
	salary int
	experience int
}

func main () {

	//var Thanga FSS

	Thanga := FSS{"YES","NO","YES",50000,3}
	fmt.Println(Thanga)

	Srivaishnavi := FSS{
		"YES",
		"YES",
		"NO",
		70000,
		3}
	fmt.Println(Srivaishnavi)


}
