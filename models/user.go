package models

import (
	"fmt"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	CreatedAt string `json:"created_at"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (u *User) Bind(r *http.Request) error {
	if u.Firstname == "" {
		return fmt.Errorf("firstname is a required field")
	}
	if u.Lastname == "" {
		return fmt.Errorf("lastname is a required field")
	}
	return nil
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
