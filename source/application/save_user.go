package application

import (
	"log"
	"wallet/source/domain"

	"github.com/google/uuid"
)

type SaveUserService struct {
}

func (s *SaveUserService) Execute(id uuid.UUID, email string, password string) error {
	user := domain.User{
		Id:       id,
		Email:    email,
		Password: password,
	}

	log.Println("user saved:", user)

	return nil
}
