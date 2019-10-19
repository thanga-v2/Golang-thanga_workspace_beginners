package main

import "fmt"



func main()  {
	a := 10

	addr_a := &a
	val_a := a

	//pointer
	//pointer *type
	var pointer_a *int

	pointer_a = &a

	fmt.Println("pointer of a",pointer_a)
	fmt.Println("value present in the pointer of a",*pointer_a)

	fmt.Println("addrees of a",addr_a,"value of a",val_a)

	//normal array

	var y_array [10]int

	fmt.Println(y_array)

	//dynamic array
	x_array := [5]int{1,2,3,4,5}
	fmt.Println(x_array)


	var my2darray [3][3]int
	my2darray[0][2]=2
	fmt.Println("my 2d array ",my2darray)


	a = len(x_array)
	//fmt.Println(a)

	for i:=0;i<len(x_array);i++ {
		fmt.Println(x_array[i])
	}

}
