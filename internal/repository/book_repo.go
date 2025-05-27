package repository

import (
"errors"
"fmt"
	"github.com/lawsonredeye/lms/internal/domain"
) 

type BookRepositoryInterface interface {
	CreateBook(title, author string, publishedYear int, genre string) string
	DeleteBookByID(id string) error
	PrintBooks()
}

type BookRepository struct {
	books map[string]*domain.Book
}

func NewBookDatabase() *BookRepository {
	return &BookRepository{
		books: make(map[string]*domain.Book),
	}
}

func (r *BookRepository) CreateBook(title, author string, publishedYear int, genre string) string {
	book := domain.NewBook(title, author, publishedYear, genre)
	r.books[book.ID] = book
	return book.ID
}

func (r *BookRepository) DeleteBookByID(id string) error {
	if ok := r.books[id]; ok == nil {
		return errors.New("invalid book id")
	}
	delete(r.books, id)
	return nil
}

func (r *BookRepository) PrintBooks() {
	for _, val := range r.books {
		fmt.Println(val.Title, val.Author, val.PublishedYear)
	}
}