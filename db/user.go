package db

import (
	"database/sql"

	"github.com/xpero39/simple-books-api/models"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}

	rows, err := db.Conn.Query("SELECT * FROM users ORDER BY ID DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) AddUser(user *models.User) error {
	var id int
	var createdAt string
	query := `INSERT INTO users (firstname, lastname) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, user.Firstname, user.Lastname).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	user.ID = id
	user.CreatedAt = createdAt
	return nil
}

func (db Database) GetUserById(userId int) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE id = $1;`
	row := db.Conn.QueryRow(query, userId)
	switch err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.CreatedAt); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}
