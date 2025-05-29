package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/lawsonredeye/lms/internal/domain"
)

type BookRepository struct {
	books      map[string]*domain.Book
	bookLedger map[string]*domain.BookLedger
}

func NewBookDatabase() *BookRepository {
	return &BookRepository{
		books:      make(map[string]*domain.Book),
		bookLedger: make(map[string]*domain.BookLedger),
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

func (r *BookRepository) GetBookByID(id string) (*domain.Book, error) {
	d, exists := r.books[id]
	if !exists {
		return nil, errors.New("book with id not found")
	}
	return d, nil
}

func (r *BookRepository) UpdateBookByID(id string, title, author string, publishedYear int, genre string) error {
	book, exists := r.books[id]
	if !exists {
		return errors.New("book not found")
	}

	if title != "" {
		book.Title = title
	}
	if author != "" {
		book.Author = author
	}
	if publishedYear > 0 {
		book.PublishedYear = publishedYear
	}
	if genre != "" {
		book.Genre = genre
	}
	book.UpdatedAt = time.Now()

	r.books[id] = book
	return nil
}

func (r *BookRepository) BorrowBook(id, lenderId string) error {
	if exists := r.books[id]; exists == nil {
		return errors.New(domain.ERRBOOKNOTFOUND)
	}

	book := domain.NewBorrowedBook(id, lenderId)
	r.bookLedger[id] = book
	return nil
}

func (r *BookRepository) ReturnedBookByID(id string) error {
	book := r.bookLedger[id]
	if book == nil {
		return errors.New(domain.ERRBOOKNOTFOUND)
	}

	book.Status = domain.RETURNED
	return nil
}

func (r *BookRepository) GetBorrowedBooks() {
	for _, val := range r.bookLedger {
		fmt.Printf("Book id: %v, lender: %v, status: %v, borrowed time: %v\n", val.BookID, val.LenderID, val.Status, val.CreatedAt)
	}
}
