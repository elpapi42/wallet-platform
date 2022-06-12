package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Email    string
	Password string
}

type CreateUserCommand struct {
	User
}

func (c *CreateUserCommand) GetKey() string {
	return c.Id.String()
}

func (c *CreateUserCommand) GetName() string {
	return "CreateUser"
}
