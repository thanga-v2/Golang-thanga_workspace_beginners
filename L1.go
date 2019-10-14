// Entry point of the program
package main

import "fmt"
// exit of the program
func main()  {
	fmt.Println("// this is to check if it is working fine ! // ")
	sec()
}
func  sec()  {
	fmt.Println("// i am executing out side //")

	for i := 0 ; i < 20 ; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
	}


	}

	fmt.Println("your iteration is  over !!")
}

////control flow (how the computer reads code)
//(1) Sequence
//(2) Iterative
//(3) Conditional
