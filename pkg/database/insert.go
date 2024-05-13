package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func AddBook(book BookDetail) error {

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/library")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	_, err = db.Exec(`INSERT INTO bookrecords(Name,Author,Publication_Year) VALUES(?,?,?)`, book.Name, book.Author, book.Publication_Year)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
