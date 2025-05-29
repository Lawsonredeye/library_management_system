package repository

import "github.com/lawsonredeye/lms/internal/domain"

type BookRepositoryInterface interface {
	CreateBook(title, author string, publishedYear int, genre string) string
	DeleteBookByID(id string) error
	PrintBooks()
	UpdateBookByID(id string, title, author string, publishedYear int, genre string) error
	GetBookByID(id string) (*domain.Book, error)
	BorrowBook(id, lenderId string) error
	ReturnedBookByID(id string) error
	GetBorrowedBooks()
}

type MemberRepositoryInterface interface {
	CreateMember(name, password string) string
	GetAllMembers() []*domain.Members
	DeleteMemberByID(id string) error
	UpdateMemberByID(id, name, password string) error
}
