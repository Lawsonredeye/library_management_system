package repository

import "github.com/lawsonredeye/lms/internal/domain"

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
