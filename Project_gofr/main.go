// main.go
package main

import (
	"project_gofr/db"
	"project_gofr/handlers"

	"gofr.dev/pkg/gofr"
)

func main() {
	// create the application object
	app := gofr.New()
	db.InitDB()
	app.Server.ValidateHeaders = false
	// add a handler
	app.GET("/books", handlers.GetBooks)

	// handler can access the parameters from context.
	app.GET("/books/{id}", handlers.GetBook)

	// handler function can send response in JSON
	app.POST("/books", handlers.AddBook)

	// handler returns response based on PathParam
	app.PUT("/books/{id}", handlers.UpdateBook)

	// Handler function which throws error
	app.DELETE("/books/{id}", handlers.DeleteBook)

	// Handler function which uses logging
	app.GET("/users", handlers.GetUsers)

	// Handler function return response of type raw
	app.GET("/users/{id}", handlers.GetUser)
	app.POST("/users", handlers.AddUser)
	app.DELETE("/users/{id}", handlers.DeleteUser)
	app.PUT("/users/{id}", handlers.UpdateUser)
	app.Start()
}
