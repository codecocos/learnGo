package main

import "fmt"

//a와 b의 type이 같다면 (a,b int) 로 쓸 수 있다.
func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}
