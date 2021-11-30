package main

import "fmt"

func canIDrink(age int) bool {
	
	//방법 1.
	// switch age {
	// 	case 10:
	// 		return false
	// 	case 18:
	// 		return true
	// }

	//방법 2.
	// switch  {
	// case age < 18 :	
	// 	return false
	// case age == 18:
	// 	return true
	// case age > 50:
	// 	return false
	// }
	// return false

	//방법 3.
	switch koreanAge := age + 2; koreanAge {
		case 10 :	
			return false
		case 18 :
			return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(18))
}