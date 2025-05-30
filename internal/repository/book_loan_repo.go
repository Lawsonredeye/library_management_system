package repository

import (
	"errors"
	"fmt"

	"github.com/lawsonredeye/lms/internal/domain"
)

type LoanRepoDB struct {
	database map[string]*domain.BookLoan
}

func NewLoanDB() *LoanRepoDB {
	return &LoanRepoDB{
		database: make(map[string]*domain.BookLoan),
	}
}

func (r *LoanRepoDB) BorrowBook(id, lenderId string) error {
	book := domain.NewBorrowedBook(id, lenderId)
	r.database[id] = book
	return nil
}

func (r *LoanRepoDB) ReturnedBookByID(id string) error {
	book, exists := r.database[id]
	if !exists {
		return errors.New(domain.ERRBOOKNOTFOUND)
	}
	book.Status = domain.RETURNED
	return nil
}

func (r *LoanRepoDB) GetBorrowedBooks() {
	for _, val := range r.database {
		if val.Status == domain.BORROWED {
			fmt.Printf("book id: %v, lender id: %v, status: %v, borrowed date: %v\n", val.BookID, val.LenderID, val.Status, val.CreatedAt)
		}
	}
}

func (r *LoanRepoDB) GetReturnedBooks() {
	for _, val := range r.database {
		if val.Status == domain.RETURNED {
			fmt.Printf("book id: %v, lender id: %v, status: %v, returned date: %v\n", val.BookID, val.LenderID, val.Status, val.UpdatedAt)
		}
	}
}
