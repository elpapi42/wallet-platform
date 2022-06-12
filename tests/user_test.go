package tests

import (
	"testing"

	"wallet/source/domain"

	"github.com/google/uuid"
)

func TestInstanceUser(t *testing.T) {
	id := uuid.New()

	user := domain.User{
		Id:       id,
		Email:    "whitman@bohorquez.com",
		Password: "123456",
	}

	if user.Id != id {
		t.Errorf("Expected %v, got %v", id, user.Id)
	}
}
