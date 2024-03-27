package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/xpero39/simple-books-api/db"
	"github.com/xpero39/simple-books-api/models"
)

var bookIDKey = "bookID"

func books(router chi.Router) {
	router.Get("/books", getAllBooks)
	router.Post("/books/create", createBook)

	router.Route("/books/{bookId}", func(router chi.Router) {
		router.Use(BookContext)
		router.Get("/", getBook)
		router.Delete("/", deleteBook)
		/* 		router.Put("/", updateBook) */
	})

	router.Route("/{bookId}", func(router chi.Router) { // is this neccessary? books/{bookId} already exists
		router.Use(BookContext)
		router.Get("/", getBook)
		router.Delete("/", deleteBook)
		/* 		router.Put("/", updateBook)
		 */
	})
}

func BookContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "bookId")
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("book ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid book ID")))
		}
		ctx := context.WithValue(r.Context(), bookIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dbInstance.GetAllBooks()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, books); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := render.Bind(r, book); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddBook(book); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, book); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.Context().Value(bookIDKey).(int)
	book, err := dbInstance.GetBookById(bookID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &book); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.Context().Value(bookIDKey).(int)
	err := dbInstance.DeleteBook(bookId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.Context().Value(bookIDKey).(int)
	bookData := models.Book{}
	if err := render.Bind(r, &bookData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	book, err := dbInstance.UpdateBook(bookId, bookData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &book); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
