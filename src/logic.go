package src

type Book struct {
	name        string
	description string
	id          int
	author      string
	genre       []string
}

type Library struct {
	books   []Book
	address string
	id      int
	members []string // need to add member struct
}

func (lib *Library) AddBooks(book []Book) {
	// this function add books list of books to library
}

func (lib *Library) RemoveBook(book Book) {
	// this function remove book from books to library
}
func (lib *Library) ShowBooks() {
	// show list of library books
}

func (lib *Library) GetBookById(book Book) {
	// get specific book by id
}
