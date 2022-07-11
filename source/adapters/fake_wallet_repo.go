package adapters

import (
	"wallet/source/domain"

	"github.com/google/uuid"
)

type FakeWalletRepository struct {
	wallets map[uuid.UUID]*domain.Wallet
}

func NewFakeWalletRepository() *FakeWalletRepository {
	return &FakeWalletRepository{
		wallets: make(map[uuid.UUID]*domain.Wallet),
	}
}

func (r *FakeWalletRepository) Add(wallet *domain.Wallet) error {
	r.wallets[wallet.GetId()] = wallet
	return nil
}

func (r *FakeWalletRepository) Get(id uuid.UUID) (*domain.Wallet, error) {
	return r.wallets[id], nil
}
