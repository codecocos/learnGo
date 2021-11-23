package main

import (
	"fmt"
	"strings"
)

//naked 함수
func lenAndUpper(name string) (length int, uppercase string) {

	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)

	//return length, uppercase : 아래와 같은 결과 값.
	return
}

func main() {

	totalLength, up := lenAndUpper("coco")
	fmt.Println(totalLength, up) // 4 , COCO

}
