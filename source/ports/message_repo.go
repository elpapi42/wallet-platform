package ports

import "wallet/source/domain/messages"

type MessageRepository interface {
	Add([]messages.Message) error
}
