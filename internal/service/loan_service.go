package service

import "github.com/lawsonredeye/lms/internal/repository"

type LoanService struct {
	service repository.LoanRepo
}

func NewLoanService(repo repository.LoanRepo) *LoanService {
	return &LoanService{
		service: repo,
	}
}

func (s *LoanService) BorrowAnyBook(id, lenderId string) error {
	return s.service.BorrowBook(id, lenderId)
}

func (s *LoanService) ReturnBookByID(id string) error {
	return s.service.ReturnedBookByID(id)
}

func (s *LoanService) PrintBorrowedBooks() {
	s.service.GetBorrowedBooks()
}

func (s *LoanService) PrintReturnedBooks() {
	s.service.GetReturnedBooks()
}
