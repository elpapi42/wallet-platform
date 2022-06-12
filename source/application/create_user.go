package application

import (
	"wallet/source/domain"
	"wallet/source/ports"

	"github.com/google/uuid"
)

type CreateUserService struct {
	MessageRepo ports.MessageRepository
}

func (s *CreateUserService) Execute(email string, password string) error {
	user := domain.User{
		Id:       uuid.New(),
		Email:    email,
		Password: password,
	}

	command := &domain.CreateUserCommand{User: user}

	error := s.MessageRepo.Add(command)
	if error != nil {
		return error
	}

	return nil
}
