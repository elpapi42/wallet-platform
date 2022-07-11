package messages

import "github.com/google/uuid"

type DepositReceived struct {
	User   uuid.UUID
	Wallet uuid.UUID
	Amount float64
}

func (c *DepositReceived) GetName() string {
	return "DepositReceived"
}

func (c *DepositReceived) GetKey() string {
	return c.User.String()
}
