package main

import "fmt"

func main ()  {

	//multiple initialisation
	// var x,y int32 = 15,16

	//decleration and initiation inferred

	//var a int32
        //or
	//var a = 15
		//or
	//a = 15
		//or

	a := 15  //allocation,initialisation,assignment all in one shot

	fmt.Println("value of a",a)

	//accessing the memory address of the variable a

	fmt.Println("memory address of a",&a)

	//var b = 15
	b:=17

	fmt.Println("value of b",b)


}


//static declaration
// var x int
// x = 100

//dynamic declaration
// x := 100


//mixed decleration
// var x,y,z = 1,6,"hi thanga"