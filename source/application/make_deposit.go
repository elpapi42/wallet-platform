package application

import (
	"log"
	"wallet/source/ports"

	"github.com/google/uuid"
)

type MakeDepositService struct {
	WalletRepository ports.WalletRepository
}

func (s *MakeDepositService) Execute(walletId uuid.UUID, amount float64) error {
	wallet, err := s.WalletRepository.Get(walletId)
	if err != nil {
		log.Println("error getting wallet:", err)
		return err
	}

	wallet.Deposit(amount)

	if err := s.WalletRepository.Add(wallet); err != nil {
		log.Println("error adding wallet to repository:", err)
		return err
	}

	log.Println("deposit received in wallet:", wallet.GetId(), wallet.GetCurrency(), wallet.GetBalance())

	return nil
}
