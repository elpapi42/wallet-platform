package domain

import (
	"github.com/google/uuid"

	"wallet/source/domain/messages"
)

// Container for Users money
type Wallet struct {
	id       uuid.UUID
	userId   uuid.UUID
	currency string
	balance  float64
	events   []messages.Message
}

func NewWallet(userId uuid.UUID, currency string) *Wallet {
	w := &Wallet{
		id:       uuid.New(),
		userId:   userId,
		currency: currency,
	}

	w.raise(&messages.WalletCreated{
		Id:       w.id,
		User:     w.userId,
		Currency: w.currency,
	})

	return w
}

func LoadWallet(id uuid.UUID, userId uuid.UUID, currency string, balance float64) *Wallet {
	w := &Wallet{
		id:       id,
		userId:   userId,
		currency: currency,
		balance:  balance,
	}

	return w
}

func NewFromEvents(events []messages.Message) *Wallet {
	w := &Wallet{}

	for _, event := range events {
		w.On(event)
	}

	return w
}

func (w *Wallet) Deposit(amount float64) {
	w.raise(&messages.DepositReceived{
		User:   w.userId,
		Wallet: w.id,
		Amount: amount,
	})
}

func (w *Wallet) On(event messages.Message) {
	switch e := event.(type) {
	case *messages.WalletCreated:
		w.id = e.Id
		w.userId = e.User
		w.currency = e.Currency
	case *messages.DepositReceived:
		w.balance += e.Amount
	}
}

func (w *Wallet) raise(event messages.Message) {
	w.events = append(w.events, event)
	w.On(event)
}

func (w *Wallet) GetId() uuid.UUID {
	return w.id
}

func (w *Wallet) GetUserId() uuid.UUID {
	return w.userId
}

func (w *Wallet) GetCurrency() string {
	return w.currency
}

func (w *Wallet) GetBalance() float64 {
	return w.balance
}

func (w *Wallet) GetEvents() []messages.Message {
	return w.events
}
