package ports

import "wallet/source/domain"

type UserRepository interface {
	Add(domain.User) error
}
