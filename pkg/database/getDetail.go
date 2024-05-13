package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetBookDetails() []BookDetail {
	var books []BookDetail

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/library")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	result, err := db.Query("SELECT * FROM bookrecords")

	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		var book BookDetail
		err = result.Scan(&book.Name, &book.Author, &book.Publication_Year)
		if err != nil {
			fmt.Println(err)
		}
		books = append(books, book)

	}
	return books

}
