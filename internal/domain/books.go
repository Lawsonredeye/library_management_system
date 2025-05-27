package domain

import (
	"time"

	"github.com/lawsonredeye/lms/pkg"
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
