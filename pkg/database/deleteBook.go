package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteBookByBookName(book_name string) error {

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/library")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	_, err = db.Exec(`DELETE * FROM bookrecords WHERE Name = ?`, book_name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
