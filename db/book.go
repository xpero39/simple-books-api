package db

import (
	"database/sql"

	"github.com/xpero39/simple-books-api/models"
)

func (db Database) GetAllBooks() (*models.BookList, error) {
	list := &models.BookList{}

	rows, err := db.Conn.Query("SELECT * FROM books ORDER BY ID DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Quantity, &book.RentedQuantity, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return list, err
		}
		list.Books = append(list.Books, book)
	}
	return list, nil
}

func (db Database) AddBook(book *models.Book) error {
	var id int
	var createdAt string
	query := `INSERT INTO books (title, quantity) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, book.Title, book.Quantity).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	book.ID = id
	book.CreatedAt = createdAt
	return nil
}

func (db Database) GetBookById(bookId int) (models.Book, error) {
	book := models.Book{}

	query := `SELECT * FROM books WHERE id = $1;`
	row := db.Conn.QueryRow(query, bookId)
	switch err := row.Scan(&book.ID, &book.Title, &book.Quantity, &book.RentedQuantity, &book.CreatedAt, &book.UpdatedAt); err {
	case sql.ErrNoRows:
		return book, ErrNoMatch
	default:
		return book, err
	}
}

// TEST this function!
func (db Database) DeleteBook(bookId int) error {
	query := `DELETE FROM books WHERE id = $1;`
	_, err := db.Conn.Exec(query, bookId)
	if err != nil {
		return err
	}
	return nil
}

// TEST this function!
func (db Database) UpdateBook(bookId int, book models.Book) (models.Book, error) {
	//Update and return the updated book
	query := `UPDATE books SET title = $1, quantity = $2, rented_quantity = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4;`
	_, err := db.Conn.Exec(query, book.Title, book.Quantity, book.RentedQuantity, bookId)

	if err != nil {
		return book, err
	}
	return book, nil
}
