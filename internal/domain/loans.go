package domain

import "time"

type BookLoan struct {
	BookID    string
	LenderID  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBorrowedBook(id, lenderId string) *BookLoan {
	now := time.Now()
	return &BookLoan{
		BookID:    id,
		LenderID:  lenderId,
		Status:    BORROWED,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
