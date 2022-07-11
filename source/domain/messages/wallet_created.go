package messages

import "github.com/google/uuid"

type WalletCreated struct {
	Id       uuid.UUID
	User     uuid.UUID
	Currency string
}

func (c *WalletCreated) GetName() string {
	return "WalletCreated"
}

func (c *WalletCreated) GetKey() string {
	return c.Id.String()
}
