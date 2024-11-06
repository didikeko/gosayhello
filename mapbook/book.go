package mapbook

import "fmt"

//create function with map
func Book() {
	book :=  make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko Kurniawan Khannedy"
	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])
}