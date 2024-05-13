package controller

import (
	"library/pkg/database"
	"library/services/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func AllBookHandler(c *gin.Context) {

	// regularBooks, adminBooks := books.GetAllBooks()

	books := database.GetBookDetails()

	// resp := AllBooksResponse{
	// 	RegularBooks: regularBooks,
	// 	AdminBooks:   adminBooks,
	// }

	c.JSON(http.StatusAccepted, books)
}

func AddBookHandler(c *gin.Context) {
	var req AddBookRequest
	var resp AddBookResponse
	err := c.BindJSON(&req)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// newBook := &csv.Book{
	// 	Name:            req.Name,
	// 	Author:          req.Author,
	// 	PublicationYear: req.PublicationYear,
	// }

	newBook := database.BookDetail{
		Name:             req.Name,
		Author:           req.Author,
		Publication_Year: req.PublicationYear,
	}
	err = books.AddBookDetails(newBook)
	//fmt.Println(err)
	// err = books.AddBook(newBook)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Message = "Book added successfully"
	c.JSON(http.StatusAccepted, resp)
}

func DeleteBookHandler(c *gin.Context) {
	var req DeleteBookRequest
	var resp DeleteBookResponse
	err := c.BindJSON(&req)
	if err != nil {
		resp.Error = err
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err = books.DeleteBook(req.Name)
	if err != nil {
		resp.Error = err
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Message = "Book deleted successfully"
	c.JSON(http.StatusAccepted, resp)
}
