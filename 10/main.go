package main

import (
	"fmt"

	"github.com/learnGo/10/mydict"
)

func main() {

	/* Search */
	// dictionary := mydict.Dictionary{"first": "First word"}

	// definition, err := dictionary.Search("first")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)

	////////////////////////////////////////////////////////////////

	/* Add */
	// dictionary := mydict.Dictionary{}

	// word := "hello"
	// definition := "Greeting"

	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, _ := dictionary.Search(word)
	// fmt.Println("found : ", word, "definition : ", hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	////////////////////////////////////////////////////////////////

	/* Update */
	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, " First")
	// err := dictionary.Update("baseWord", "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	/* Delete */
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, " First")
	
	dictionary.Search(baseWord)
  //Delete는 반환하는 것이 없음.
	dictionary.Delete(baseWord)

	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)

}