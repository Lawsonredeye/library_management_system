package service

import (
	"errors"

	"github.com/lawsonredeye/lms/internal/domain"
	"github.com/lawsonredeye/lms/internal/repository"
)

type BookService struct {
	// BookRepo *repository.BookRepository
	BookRepo repository.BookRepositoryInterface
}

func NewBookService(repo repository.BookRepositoryInterface) *BookService {
	return &BookService{
		BookRepo: repo,
	}
}

func (b *BookService) CreateBook(title, author string, publishedYear int, genre string) string {
	id := b.BookRepo.CreateBook(title, author, publishedYear, genre)
	return id
}

func (b *BookService) DeleteBookByID(id string) error {
	return b.BookRepo.DeleteBookByID(id)
}

func (b *BookService) PrintBooks() {
	b.BookRepo.PrintBooks()
}

func (b *BookService) UpdateBookByID(id string, title, author string, publishedYear int, genre string) error {
	return b.BookRepo.UpdateBookByID(id, title, author, publishedYear, genre)
}

func (b *BookService) GetBookByID(id string) (*domain.Book, error) {
	data, err := b.BookRepo.GetBookByID(id)
	if err != nil {
		return nil, errors.New("book with id not found")
	}
	return data, nil
}
