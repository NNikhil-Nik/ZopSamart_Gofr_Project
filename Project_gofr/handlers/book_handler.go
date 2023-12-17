// handlers/book_handler.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"project_gofr/db"
	"project_gofr/models"
	"strconv"

	"gofr.dev/pkg/gofr"
)

// GetBooksHandler is a gofr.Handler that returns a list of all books in the library
func GetBooksHandler(ctx *gofr.Context) (interface{}, error) {
	books, err := db.GetBooks()
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return books, nil
}

// GetBookHandler is a gofr.Handler that returns a specific book by ID
func GetBookHandler(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	book, err := db.GetBook(id)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusNotFound, "Not Found")
	}

	return book, nil
}

// AddBookHandler is a gofr.Handler that adds a new book to the library
func AddBookHandler(ctx *gofr.Context) (interface{}, error) {
	var book models.Book
	err := json.NewDecoder(ctx.Request.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	// Validate the book data (you can add more validation as needed)

	// Add the book to the database
	err = db.AddBook(&book)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return book, nil
}

// UpdateBookHandler is a gofr.Handler that updates the information of a book
func UpdateBookHandler(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	var book models.Book
	err = json.NewDecoder(ctx.Request.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	// Validate the book data (you can add more validation as needed)

	// Update the book in the database
	err = db.UpdateBook(id, &book)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return book, nil
}

// DeleteBookHandler is a gofr.Handler that removes a book from the library
func DeleteBookHandler(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	// Delete the book from the database
	err = db.DeleteBook(id)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil, nil
}
