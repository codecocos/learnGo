package main

import "fmt"

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("coco", "stone", "black", "dal", "perfume")
}
