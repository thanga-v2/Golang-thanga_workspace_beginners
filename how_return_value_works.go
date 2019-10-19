package main

import "fmt"

func find_biggest(q int,w int) int{
	if q > w {
		//fmt.Println(q)
		return q
	} else {
		//fmt.Println(w)
		return w
	}




}

func main()  {
	find_biggest(23,55)

	fmt.Println(find_biggest(23, 55))

	//fmt.Println(find_biggest(&a int,&b int()))



}