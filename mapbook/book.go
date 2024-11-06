package mapbook

//create function with map
func Book() map[string]string {
	book :=  make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko Kurniawan Khannedy"
	return book
}