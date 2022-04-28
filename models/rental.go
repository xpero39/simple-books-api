package models

import (
	"fmt"
	"net/http"
)

type Rental struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	BookId    int    `json:"book_id"`
	Quantity  int    `json:"quantity"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RentalList struct {
	Rentals []Rental `json:"rentals"`
}

func (rt *Rental) Bind(r *http.Request) error {
	if rt.Quantity == 0 {
		return fmt.Errorf("rented quantity is a required field")
	}
	return nil
}

func (*RentalList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Rental) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
