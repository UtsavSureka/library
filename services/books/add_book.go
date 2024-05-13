package books

import (
	"errors"
	"library/pkg/database"
)

// func AddBook(newBook *csv.Book) error {

// 	// validate if the newBook exists
// 	for _, book := range csv.Get(constants.RegularUserCSV) {
// 		if book.Name == newBook.Name {
// 			return errors.New("book already exists")
// 		}
// 	}
// 	csv.Append(constants.RegularUserCSV, newBook)
// 	return nil
// }

func AddBookDetails(newBook database.BookDetail) error {
	for _, book := range database.GetBookDetails() {
		if book.Name == newBook.Name {
			return errors.New("book already exists")
		}
	}
	database.AddBook(newBook)
	return nil
}
