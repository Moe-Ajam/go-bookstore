package bookstore

import "errors"

type Book struct {
	Id     int
	Title  string
	Author string
	Copies int
}

func Buy(book Book) (Book, error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	book.Copies--
	return book, nil
}

func GetAllBooks(books []Book) []Book {
	return books
}

//func GetBook(books []Book, id int) Book {
//	for _, b := range books {
//		if b.Id == id {
//			return b
//		}
//	}
//}
