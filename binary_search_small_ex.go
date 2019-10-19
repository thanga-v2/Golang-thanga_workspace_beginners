package main

import "fmt"

func checkBin(list []string, i string) bool {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == i {
			return true
		}
		if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(checkBin([]string{"thanga", "barath", "vaishu", "amma", "siva"}, "siva"))
}
