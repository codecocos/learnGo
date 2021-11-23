package main1

import "fmt"

func superAdd(numbers ...int) int {
	// fmt.Println(numbers) // [ 1 2 3 4 5 6]
	
	//for : numbers 범위 안에서 조건에 따라 반복 실행함
	//range : index를 부여, for에서만 사용가능
	for index, number := range numbers {
		fmt.Println(index, number) // 0 1 2 3 4 5 \n 1 2 3 4 5 6
	}
	return 1
}

func superAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++{
		fmt.Println((numbers[i]))// 1 2 3 4 5 6
	}
	return 1
}

func main1() {
	superAdd(1, 2, 3, 4, 5, 6)
	superAdd2(1, 2, 3, 4, 5, 6)
}