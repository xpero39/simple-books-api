package db

import (
	"database/sql"

	"github.com/xpero39/simple-books-api/models"
)

func (db Database) GetAllRentals() (*models.RentalList, error) {
	list := &models.RentalList{}

	rows, err := db.Conn.Query("SELECT * FROM rentals ORDER BY ID DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var rental models.Rental
		err := rows.Scan(&rental.ID)
		if err != nil {
			return list, err
		}
		list.Rentals = append(list.Rentals, rental)
	}
	return list, nil
}

func (db Database) AddRental(rental *models.Rental) error {
	var id int
	var createdAt string

	//Get book and user id

	query := `INSERT INTO rentals (book_id, user_id, quantity) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, rental.BookId, rental.UserId, rental.Quantity).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	rental.ID = id
	rental.CreatedAt = createdAt
	return nil
}

func (db Database) CloseRental(bookId int, userId int, qty int) error {
	//Find

	return nil
}

func (db Database) GetRentalById(rentalId int) (models.Rental, error) {
	rental := models.Rental{}

	query := `SELECT * FROM rentals WHERE id = $1;`
	row := db.Conn.QueryRow(query, rentalId)
	switch err := row.Scan(&rental.ID, &rental.BookId, &rental.UserId, &rental.Quantity, &rental.CreatedAt, &rental.UpdatedAt); err {
	case sql.ErrNoRows:
		return rental, ErrNoMatch
	default:
		return rental, err
	}
}
