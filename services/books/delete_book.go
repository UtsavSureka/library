package books

import (
	"errors"
	"library/pkg/database"
)

func DeleteBook(bookName string) error {

	// validate if the book is present or not

	//allBooks := csv.Get(constants.RegularUserCSV)

	allBooks := database.GetBookDetails()

	//newBooks := []csv.Book{}
	//bookFound := false

	bookFound := false

	// for _, book := range allBooks {
	// 	if book.Name != bookName {
	// 		newBooks = append(newBooks, book)
	// 	} else {
	// 		bookFound = true
	// 	}
	// }

	for _, book := range allBooks {
		if book.Name == bookName {
			bookFound = true
		}
	}
	if !bookFound {
		return errors.New("book not found")
	}

	//csv.Rewrite(constants.RegularUserCSV, newBooks)
	database.DeleteBookByBookName(bookName)
	return nil
}
