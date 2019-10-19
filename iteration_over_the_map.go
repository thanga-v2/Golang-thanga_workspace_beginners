package main

import "fmt"

func main() {

	//range helps to iterate over the map
   var school_slam = map[string]int{
   	"thanga":25,
   	"barath":24,
   	"siva":22,
   	"vaishu":25,
   }


   for name,age := range school_slam {
   	fmt.Println(name,age)
   }
}
