package repository

import (
	"errors"
	"time"

	"github.com/lawsonredeye/lms/internal/domain"
)

const (
	ERRIDNOTFOUND string = "member with id not found"
)

type MemberRepoDB struct {
	database map[string]*domain.Members
}

func NewMemberDatabase() *MemberRepoDB {
	return &MemberRepoDB{
		database: make(map[string]*domain.Members),
	}
}

func (r *MemberRepoDB) CreateMember(name, password string) string {
	member := domain.NewMember(name, password)
	r.database[member.ID] = member
	return member.ID
}

func (r *MemberRepoDB) GetAllMembers() []*domain.Members {
	resp := make([]*domain.Members, 0, len(r.database))

	for _, val := range r.database {
		resp = append(resp, val)
	}
	return resp
}

func (r *MemberRepoDB) DeleteMemberByID(id string) error {
	m := r.database[id]
	if m == nil {
		return errors.New(ERRIDNOTFOUND)
	}

	delete(r.database, m.ID)
	return nil
}

func (r *MemberRepoDB) UpdateMemberByID(id, name, password string) error {
	m := r.database[id]
	if m == nil {
		return errors.New(ERRIDNOTFOUND)
	}

	if name != "" {
		m.Name = name
	}
	if password != "" && len(password) > 7 {
		m.Password = password
	}

	m.UpdatedAt = time.Now()
	return nil
}
