package ports

import "wallet/source/domain"

type MessageRepository interface {
	Add(domain.Message) error
}
