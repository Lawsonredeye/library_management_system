package service

import "github.com/lawsonredeye/lms/internal/repository"

type BookService struct {
	BookRepo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
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
