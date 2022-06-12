package domain

type Message interface {
	GetKey() string
	GetName() string
}
