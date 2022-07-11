package application

import (
	"log"
	"wallet/source/domain"
	"wallet/source/ports"

	"github.com/google/uuid"
)

type CreateWalletService struct {
	WalletRepository ports.WalletRepository
}

func (s *CreateWalletService) Execute(userId uuid.UUID, currency string) (*domain.Wallet, error) {
	wallet := domain.NewWallet(userId, currency)

	if error := s.WalletRepository.Add(wallet); error != nil {
		return nil, error
	}

	log.Println("wallet created:", wallet.GetId(), wallet.GetUserId(), wallet.GetCurrency(), wallet.GetBalance())

	return wallet, nil
}
