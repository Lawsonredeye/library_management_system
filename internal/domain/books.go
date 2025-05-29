package domain

import (
	"time"

	"github.com/lawsonredeye/lms/pkg"
)

const (
	BORROWED        string = "borrowed"
	RETURNED        string = "returned"
	ERRBOOKNOTFOUND string = "book with id not found"
)

type Book struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedYear int       `json:"published_year"`
	Genre         string    `json:"genre"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BookLedger struct {
	BookID    string
	LenderID  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBook(title, author string, publishedYear int, genre string) *Book {
	id, _ := pkg.GenerateUUID()
	now := time.Now()
	return &Book{
		ID:            id,
		Title:         title,
		Author:        author,
		PublishedYear: publishedYear,
		Genre:         genre,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

func NewBorrowedBook(id, lenderId string) *BookLedger {
	now := time.Now()
	return &BookLedger{
		BookID:    id,
		LenderID:  lenderId,
		Status:    BORROWED,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
