// db/queries.go
package db

import (
	"log"
	"project_gofr/models"
)

// GetBooks retrieves all books from the database
func GetBooks() ([]models.Book, error) {
	rows, err := DB.Query("SELECT id, title, author, quantity FROM books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	books := make([]models.Book, 0)
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// GetBook retrieves a specific book by ID from the database
func GetBook(id int) (*models.Book, error) {
	var book models.Book
	err := DB.QueryRow("SELECT id, title, author, quantity FROM books WHERE id = ?", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &book, nil
}

// AddBook adds a new book to the database
func AddBook(book *models.Book) error {
	result, err := DB.Exec("INSERT INTO books (title, author, quantity) VALUES (?, ?, ?)",
		book.Title, book.Author, book.Quantity)
	if err != nil {
		log.Println(err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return err
	}

	book.ID = int(id)
	return nil
}

// UpdateBook updates the information of a book in the database
func UpdateBook(id int, book *models.Book) error {
	_, err := DB.Exec("UPDATE books SET title = ?, author = ?, quantity = ? WHERE id = ?",
		book.Title, book.Author, book.Quantity, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// DeleteBook removes a book from the database
func DeleteBook(id int) error {
	_, err := DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetUsers retrieves all users from the database
func GetUsers() ([]models.User, error) {
	rows, err := DB.Query("SELECT id, name FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Fetch the books associated with the user
		user.Books, err = getUserBooks(user.ID)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUser retrieves a specific user by ID from the database
func GetUser(id int) (*models.User, error) {
	var user models.User
	err := DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Fetch the books associated with the user
	user.Books, err = getUserBooks(user.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

// AddUser adds a new user to the database
func AddUser(user *models.User) error {
	result, err := DB.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return err
	}

	user.ID = int(id)
	return nil
}

// UpdateUser updates the information of a user in the database
func UpdateUser(id int, user *models.User) error {
	_, err := DB.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// DeleteUser removes a user from the database
func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// getUserBooks retrieves the books associated with a user
func getUserBooks(userID int) ([]models.Book, error) {
	rows, err := DB.Query("SELECT b.id, b.title, b.author, b.quantity FROM user_books ub "+
		"JOIN books b ON ub.book_id = b.id WHERE ub.user_id = ?", userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	books := make([]models.Book, 0)
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
