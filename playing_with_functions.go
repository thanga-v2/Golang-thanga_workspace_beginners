package main


//functions -> write something once and re-use it everywere

//modular code

//basic functions
/*
func main(){

}

*/

/*
func name(value 1 type,value 2 type) return_type {
<code>
}

 */




/*

import (
	"fmt"
)

func c(x int32){
	//i := 10
	var i int32
	for i < x {
		fmt.Println(x)
		fmt.Println(i)
		i++
	}
}

func main()  {
	c(3)


}


 */

import (
	"fmt"
)


func reset_val(i int) {
	fmt.Println("value of i is ",i, "i is located at ",&i)
	i = 10
	fmt.Println("value of i has been modified to",i,"i is located at",&i)
}

func main()  {

	j := 20
	fmt.Println("j is located at",&j,"value of j is",j)
	reset_val(j)
	fmt.Println("j is located at",&j,"value of j is",j)
}

// we pass the data into golang by value rather than the refernce
//redecleration is possible within the function as the value gets stored into the stack