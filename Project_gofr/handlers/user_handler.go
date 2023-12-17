// handlers/user_handler.go
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

// GetUsersHandler is a gofr.Handler that returns a list of all users
func GetUsersHandler(ctx *gofr.Context) (interface{}, error) {
	users, err := db.GetUsers()
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return users, nil
}

// GetUserHandler is a gofr.Handler that returns a specific user by ID
func GetUserHandler(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	user, err := db.GetUser(id)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusNotFound, "Not Found")
	}

	return user, nil
}

// AddUserHandler is a gofr.Handler that adds a new user
func AddUserHandler(ctx *gofr.Context) (interface{}, error) {
	var user models.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	// Validate the user data (you can add more validation as needed)

	// Add the user to the database
	err = db.AddUser(&user)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return user, nil
}

// UpdateUserHandler is a gofr.Handler that updates the information of a user
func UpdateUserHandler(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	var user models.User
	err = json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	// Validate the user data (you can add more validation as needed)

	// Update the user in the database
	err = db.UpdateUser(id, &user)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return user, nil
}

// DeleteUserHandler is a gofr.Handler that removes a user
func DeleteUserHandler(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusBadRequest, "Bad Request")
	}

	// Delete the user from the database
	err = db.DeleteUser(id)
	if err != nil {
		log.Println(err)
		return nil, gofr.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil, nil
}
