package models

import (
	"fmt"
	"net/http"
)

type Book struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Quantity       int    `json:"quantity"`
	RentedQuantity int    `json:"rented_quantity"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type BookList struct {
	Books []Book `json:"books"`
}

func (b *Book) Bind(r *http.Request) error {
	if b.Title == "" {
		return fmt.Errorf("title is a required field")
	}
	return nil
}

func (*BookList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Book) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
