package main1

import (
	"fmt"
	"strings"
)

//멀티 반환 함수
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main1() {
	totalLength, upperName := lenAndUpper("coco")
	fmt.Println(totalLength, upperName) // 4 , COCO

	//totalLength, _ := lenAndUpper("coco")
	//fmt.Println(totalLength) // 4
}
