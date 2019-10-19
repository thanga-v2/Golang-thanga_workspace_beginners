package main

import (
	"fmt"
	"time"
)

func main()  {

	i := 10


	// for <init stmnt>;<exp>;<post stmnt>
	// init and post stmnt are completely optional

	for i > 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i --

	}
fmt.Println("!!!!")
}


// if <bool exp>
//<code>
//}