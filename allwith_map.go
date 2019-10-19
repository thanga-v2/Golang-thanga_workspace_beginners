
//Map is simply a key value pair

//   ------------------------   map[KeyType]ValueType   -------------------------

// 								Var x map[KeyType]valueTpe
//								var thanga [age]21


package main
import "fmt"

func main() {
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	m["one hundred"] = 100
	fmt.Println(m)
	m["age"]=21
	fmt.Println(m)
	m["add"]=11
	fmt.Println(m)
	//it will keep on adding


	var x = make(map[string]string)
	fmt.Println(x)
	//x is empty
	x["first one"]="new entry"
	fmt.Println(x)
	x["second one"]="second entry"
	fmt.Println(x)



}




