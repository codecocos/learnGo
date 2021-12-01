package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("not Found")
	errCantUpdate = errors.New("cant update non-existing word")
	errWordExists = errors.New("that word already exists")
)

//go에서는 type도 method를 가질 수 있음
// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "",errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) (error) {
	
	_, err := d.Search(word)

	
	//추가하고자 하는 단어가 없는 경우 switch문
	switch err {
		case errNotFound:
			d[word] = def
		case nil:
			return errWordExists
	}

	// //추가하고자 하는 단어가 없는 경우 if문
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	return nil
}

//Update all word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
		case nil:
			d[word] = definition
		case errNotFound:
			return errCantUpdate
	}
	return nil  
}

//(d *Dictionary)가 아닌 이유
//Dictonary is a hashmap. And by default a hashmap already has the * included.
// Delete a word
func (d Dictionary) Delete(word string) error {
	

	_, err := d.Search(word)

	switch err {
		case nil:
			delete(d, word)
		case errNotFound:
			return errNotFound
	}
	return nil  
}