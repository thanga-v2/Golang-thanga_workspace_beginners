package main

import "fmt"

func main()  {

	for today:=1;today <=7 ;today++ {
	fmt.Println("on this day",today,"cherish because ")
		switch today {
		case 1 :
			fmt.Println("week off")
			fallthrough
		case 2 :
			fmt.Println("its monday !")
			fallthrough

		case 3 :
			fmt.Println("start to work on new goals !")
			fallthrough
		case 4:
			fmt.Println("you are in the middle of the week make it productive")
			fallthrough

		case 5:
			fmt.Println("almost finishing this week")
			fallthrough

		case 6:
			fmt.Println("throw yourself a little party in case you have a successfull week")
			fallthrough
		case 7:
			fmt.Println("relax :)")
//no fallthrough after the last case




		}

		fmt.Println("")

	}

}