package service

import (
	"fmt"

	"github.com/lawsonredeye/lms/internal/domain"
	"github.com/lawsonredeye/lms/internal/repository"
)

type MemberService struct {
	repo repository.MemberRepositoryInterface
}

func NewMemberService(repo repository.MemberRepositoryInterface) *MemberService {
	return &MemberService{
		repo: repo,
	}
}

func (s *MemberService) CreateMember(name, password string) string {
	resp := s.repo.CreateMember(name, password)
	fmt.Println("member created with id:", resp)
	return resp
}

func (s *MemberService) GetAllMembers() []*domain.Members {
	resp := s.repo.GetAllMembers()
	return resp
}

func (s *MemberService) DeleteMemberByID(id string) error {
	return s.repo.DeleteMemberByID(id)
}

func (s *MemberService) UpdateMemberByID(id, name, password string) error {
	return s.repo.UpdateMemberByID(id, name, password)
}
