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
