package ports

import (
	"wallet/source/domain"

	"github.com/google/uuid"
)

type WalletRepository interface {
	Add(*domain.Wallet) error
	Get(uuid.UUID) (*domain.Wallet, error)
}
