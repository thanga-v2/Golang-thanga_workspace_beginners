package main

import "fmt"

func main()  {

	var directory = make(map[string]string)

	fmt.Println(directory)

	directory_mod := map[string]string{
		"thangaraj":"9629056516",
		"srivaishnavi":"9994059653",
		"vaishu":"7538809672",
	}
fmt.Println(directory_mod)

var thanga_mobile = directory_mod["thangaraj"]
fmt.Println("thanga mobile number is", thanga_mobile)

var vaishu_mobile = directory_mod["srivaishnavi"]
fmt.Println("vaishu number is",vaishu_mobile)

}
