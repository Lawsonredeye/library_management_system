package domain

import (
	"time"

	"github.com/lawsonredeye/lms/pkg"
)

type Members struct {
	ID        string
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMember(name, password string) *Members {
	id, _ := pkg.GenerateUUID()
	now := time.Now()
	return &Members{
		ID:        id,
		Name:      name,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
