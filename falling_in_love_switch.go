package main

import "fmt"

func main()  {

	for today:=1;today <7 ;today++ {
		switch today {
		case 1 :
			fmt.Println("week off")
		case 2 :
			fmt.Println("its monday !")

		case 3 :
			fmt.Println("start to work on new goals !")

			//and so on..
		}

	}

}
